package server

import (
	"context"
	"time"

	v1 "gitee.com/qciip-icp/v-trace/api/ca/v1"
	"gitee.com/qciip-icp/v-trace/app/ca/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGrpcServer)

func NewGrpcServer(ctx context.Context, service *service.CaService) *grpc.Server {
	opts := make([]grpc.ServerOption, 0)

	opts = append(opts, grpc.WithNetwork("tcp"))
	opts = append(opts, grpc.WithAddress("0.0.0.0:10105"))
	opts = append(opts, grpc.WithTimeout(time.Second))

	// if c.Server.Grpc.Network != "" {
	// 	opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	// }
	// if c.Server.Grpc.Address != "" {
	// 	opts = append(opts, grpc.Address(c.Server.Grpc.Address))
	// }
	// if c.Server.Grpc.Timeout != nil {
	// 	opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	// }

	srv := grpc.NewServer(opts...)
	v1.RegisterCAServiceServer(srv, service)

	return srv
}
