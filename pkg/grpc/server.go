package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/host"

	"gitee.com/qciip-icp/v-trace/pkg/endpoint"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/gogo/protobuf/jsonpb"
	metrics "github.com/grpc-ecosystem/go-grpc-middleware/providers/openmetrics/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/runtime/protoiface"
)

type GrpcConfig struct {
	Network string         `mapstructure:"network"`
	Address string         `mapstructure:"address"`
	Timeout *time.Duration `mapstructure:"timeout"`
}

// ServerOption is gRPC server option.
type ServerOption func(o *Server)

// WithNetwork with server network.
func WithNetwork(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// WithAddress with server address.
func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// WithTimeout with server timeout.
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithLogger with server logger.
func WithLogger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

// WithMiddleware with server middleware.
func WithMiddleware(m ...middleware.Middleware) ServerOption {
	return func(s *Server) {
		s.middleware = m
	}
}

// WithTLSConfig with TLS config.
func WithTLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) {
		s.tlsConf = c
	}
}

// WithListener with server lis.
func WithListener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}

// WithUnaryInterceptor returns a ServerOption that sets the UnaryServerInterceptor for the server.
func WithUnaryInterceptor(in ...grpc.UnaryServerInterceptor) ServerOption {
	return func(s *Server) {
		s.unaryInts = in
	}
}

// WithStreamInterceptor returns a ServerOption that sets the StreamServerInterceptor for the server.
func WithStreamInterceptor(in ...grpc.StreamServerInterceptor) ServerOption {
	return func(s *Server) {
		s.streamInts = in
	}
}

// WithGrpcOptions with grpc options.
func WithGrpcOptions(opts ...grpc.ServerOption) ServerOption {
	return func(s *Server) {
		s.grpcOpts = opts
	}
}

// Server is a gRPC server wrapper.
type Server struct {
	*grpc.Server
	baseCtx    context.Context
	tlsConf    *tls.Config
	lis        net.Listener
	err        error
	network    string
	address    string
	endpoint   *url.URL
	timeout    time.Duration
	log        *log.Helper
	middleware []middleware.Middleware
	unaryInts  []grpc.UnaryServerInterceptor
	streamInts []grpc.StreamServerInterceptor
	grpcOpts   []grpc.ServerOption
}

// NewServer creates a gRPC server by options.
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx: context.Background(),
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		log:     log.NewHelper(log.GetLogger()),
	}

	registry := prometheus.NewPedanticRegistry()

	unaryInts := []grpc.UnaryServerInterceptor{
		unaryServerLogInterceptor(),
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Errorf("GRPC server recovery with error: %+v", p)
			if e, ok := p.(error); ok {
				return fmt.Errorf("InternalServerError: %+v", e)
			}

			return fmt.Errorf("InternalError: recovery convert error fail")
		})),
		metrics.UnaryServerInterceptor(metrics.NewRegisteredServerMetrics(registry)),
	}
	streamInts := []grpc.StreamServerInterceptor{
		recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			logger.Errorf("GRPC server recovery with error: %+v", p)
			if e, ok := p.(error); ok {
				return fmt.Errorf("InternalServerError: %+v", e)
			}

			return fmt.Errorf("InternalError: recovery convert error fail")
		})),
	}
	for _, o := range opts {
		o(srv)
	}
	if len(srv.unaryInts) > 0 {
		unaryInts = append(unaryInts, srv.unaryInts...)
	}
	if len(srv.streamInts) > 0 {
		streamInts = append(streamInts, srv.streamInts...)
	}
	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unaryInts...),
		grpc.ChainStreamInterceptor(streamInts...),
	}
	if srv.tlsConf != nil {
		grpcOpts = append(grpcOpts, grpc.Creds(credentials.NewTLS(srv.tlsConf)))
	}
	if len(srv.grpcOpts) > 0 {
		grpcOpts = append(grpcOpts, srv.grpcOpts...)
	}
	srv.Server = grpc.NewServer(grpcOpts...)
	// listen and endpoint
	srv.err = srv.listenAndEndpoints()
	// internal register
	reflection.Register(srv.Server)
	return srv
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.baseCtx = ctx
	logger.Infof("[gRPC] server listening on: %s", s.lis.Addr().String())

	return s.Serve(s.lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {

	s.GracefulStop()

	logger.Infof("[gRPC] server stopping")

	return nil
}

func (s *Server) Endpoint() *url.URL {
	if s.endpoint != nil {
		return s.endpoint
	}

	return nil
}

func (s *Server) listenAndEndpoints() error {
	if s.lis == nil {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			return err
		}
		s.lis = lis
	}

	addr, err := host.Extract(s.address, s.lis)
	if err != nil {
		_ = s.lis.Close()
		return err
	}

	s.endpoint = endpoint.NewEndpoint("", addr, s.tlsConf != nil)

	return nil
}

var jsonPbMarshaller = &jsonpb.Marshaler{
	OrigName: true,
}

// 打印请求日志.
func unaryServerLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		method := strings.Split(info.FullMethod, "/")
		action := method[len(method)-1]
		if p, ok := req.(protoiface.MessageV1); ok {
			if content, err := jsonPbMarshaller.MarshalToString(p); err != nil {
				logger.Errorf("failed to marshal proto message to string [%s] [%+v]", action, err)
			} else {
				logger.Infof("request received [%s] [%s]", action, content)
			}
		}

		// 开启注释统计请求服务时间
		// start := time.Now()

		resp, err = handler(ctx, req)

		// elapsed := time.Since(start)
		// logger.Infof("handled request [%s] exec_time is [%s]", action, elapsed)
		// if e, ok := status.FromError(err); ok {
		// 	if e.Code() != codes.OK {
		// 		logger.Debugf("response is error: %s, %s", e.Code().String(), e.Message())
		// 	}
		// }
		return resp, err
	}
}
