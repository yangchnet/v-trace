package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	iamV1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/authz"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	localGrpc "gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/token"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/marmotedu/component-base/pkg/util/stringutil"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	Authorization = "Authorization"
	RequestIdKey  = "X-Request-Id"
	xForwardedFor = "X-Forwarded-For"
)

type Gateway struct {
	tokenMaker token.Maker

	getters map[string]localGrpc.ConnGetInterface

	logger *logger.Logger

	authz authz.AuthzServer
}

func New(ctx context.Context, r registry.Registrar, tm token.Maker, l *logger.Logger) *Gateway {
	getters := make(map[string]localGrpc.ConnGetInterface)

	// iamWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Iam)
	// if err != nil {
	// 	panic(err)
	// }
	// getters[constants.Iam] = localGrpc.NewConnGetter(ctx, iamWatcher, constants.Iam)

	vtraceWatcher, err := r.(*registry.Registry).Watch(ctx, constants.VTrace)
	if err != nil {
		panic(err)
	}
	getters[constants.VTrace] = localGrpc.NewConnGetter(ctx, vtraceWatcher, constants.VTrace)

	return &Gateway{
		tokenMaker: tm,
		getters:    getters,
		logger:     l,
		authz:      authz.NewAuthzServer(ctx),
	}
}

func (s *Gateway) Iam() iamV1.IamServiceClient {
	conn := s.getters["iam"].Get()

	return iamV1.NewIamServiceClient(conn)
}

func (s *Gateway) MainHandler(ctx context.Context) http.Handler {
	gwmux := runtime.NewServeMux(
		// gwruntime.WithErrorHandler(),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs( // FIXME
				"username", req.Header.Get("username"),
				// RequestIdKey, req.Header.Get(RequestIdKey),
			)
		}),
	)

	vtraceConn := s.getters[constants.VTrace].Get()
	if err := v1.RegisterVTraceInterfaceHandler(ctx, gwmux, vtraceConn.(*grpc.ClientConn)); err != nil {
		panic(fmt.Sprintf("error registers service to apigateway: ", err))
	}

	mux := http.NewServeMux()
	mux.Handle("/", s.serveMuxHandle(ctx, gwmux))

	return s.formWrapper(mux)
}

// serveMuxHandle set skip special route and do auth.
func (s *Gateway) serveMuxHandle(ctx context.Context, mux *runtime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if s.skipSpecialRoutes(ctx, mux, w, req) {
			// special route skip sender auth check
			mux.ServeHTTP(w, req)
		} else {
			if err := s.handleAuthorization(mux, w, req); err != nil {
				_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
				runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)

				return
			}
			mux.ServeHTTP(w, req)
		}
	})
}

// handleAuthorization 处理token并做鉴权.
func (s *Gateway) handleAuthorization(mux *runtime.ServeMux, w http.ResponseWriter, req *http.Request) error {
	var err error
	ctx := req.Context()

	auth := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
	if stringutil.FindString([]string{"Bearer"}, auth[0]) < 0 {
		return status.Errorf(codes.InvalidArgument, "http header `Authorization` format error")
	}

	payload, err := s.handleBearer(ctx, auth[1])
	if err != nil {
		return status.Error(codes.Internal, "")
	}

	urls := strings.Split(req.URL.Path, ":")
	obj := urls[0]
	act := urls[len(urls)-1]
	logger.Debugf("auth, sub: [%s], obj: [%s], act: [%s]", payload.Username, obj, act)

	cando := s.authz.CanDo(ctx, &authz.AuthzRequest{
		Sub: payload.Role,
		Obj: obj,
		Act: act,
	})
	if !cando {
		return errors.New("拒绝访问，无权限访问该接口")
	}

	req.Header.Set("username", payload.Username)
	req.Header.Del(Authorization)

	return nil
}

func (s *Gateway) handleBearer(ctx context.Context, token string) (*token.Payload, error) {
	payload, err := s.tokenMaker.VerifyToken(token)
	if err != nil {
		s.logger.Error("verify token failed: ", err)
		return nil, err
	}

	return payload, nil
}

// skipSpecialRoutes will skip the special routes without auth.
func (s *Gateway) skipSpecialRoutes(ctx context.Context, mux *runtime.ServeMux, w http.ResponseWriter, r *http.Request) bool {
	// switch r.URL.Path {
	// case "/api/v1/echo":
	// 	return true
	// case "/v1/internal/test/grpc_error":
	// 	return true
	// }

	// return false

	return match(r.URL.Path)
}

func match(path string) bool {
	re := regexp.MustCompile(`^/api/v\d+/token$|^/api/v\d+/register$|^/api/healthcheck$`)

	return re.Match([]byte(path))
}

func (s *Gateway) formWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)

				return
			}
			jsonMap := make(map[string]interface{}, len(r.Form))
			for k, v := range r.Form {
				if len(v) > 0 {
					jsonMap[k] = v[0]
				}
			}
			jsonBody, err := json.Marshal(jsonMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			r.Body = ioutil.NopCloser(bytes.NewReader(jsonBody))
			r.ContentLength = int64(len(jsonBody))
			r.Header.Set("Content-Type", "application/json")
		}

		h.ServeHTTP(w, r)
	})
}
