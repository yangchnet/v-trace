package server

import (
	"fmt"
	"io/fs"
	"mime"
	"net/http"

	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/gateway"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/server/middleware"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	third_party "gitee.com/qciip-icp/v-trace/pkg/third-party"
	"gitee.com/qciip-icp/v-trace/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"golang.org/x/net/context"
)

var HttpProvider = wire.NewSet(NewHttpServer)

type HttpServer struct {
	*gin.Engine

	conf *conf.HttpConfig

	gw *gateway.Gateway

	logger *logger.Logger
}

func NewHttpServer(ctx context.Context, c *conf.Bootstrap) *HttpServer {
	gin.SetMode(gin.DebugMode)

	r := registry.NewRegistry(ctx, &c.Etcd)

	tm, err := token.NewJWTMaker(&c.Gateway.Token)
	if err != nil {
		panic(fmt.Sprintf("failed to create token maker:%v", err))
	}

	logPrefix := "--> api-gateway ##"

	l := logger.New(logger.WithPrefix(&logPrefix))
	logger.SetLogger(&c.Gateway.Log, l)

	gw := gateway.New(ctx, r, tm, l)

	return &HttpServer{
		Engine: gin.Default(),
		conf:   &c.Gateway.Http,
		gw:     gw,
		logger: l,
	}
}

func (s *HttpServer) Start(ctx context.Context) error {
	s.Use(
		middleware.Log(ctx, s.logger),
		middleware.Recover(ctx, s.logger),
		middleware.CORSMiddleware(ctx),
	)

	s.GET("/openapi-ui/*filepath", gin.WrapH(GetOpenAPIHandler()))

	s.Any("/api/v1/*filepath", gin.WrapH(s.gw.MainHandler(ctx)))

	return s.Run(fmt.Sprintf(":%d", s.conf.Port))
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return nil
}

// GetOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func GetOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return http.StripPrefix("/openapi-ui/", http.FileServer(http.FS(subFS)))
}
