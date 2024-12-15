package server

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/conf"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGrpcServer)

func NewGrpcServer(ctx context.Context, c *conf.Bootstrap, service v1.AlgoServiceServer) *grpc.Server {
	opts := make([]grpc.ServerOption, 0)
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.WithNetwork(c.Grpc.Network))
	}
	if c.Grpc.Address != "" {
		opts = append(opts, grpc.WithAddress(c.Grpc.Address))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.WithTimeout(*c.Grpc.Timeout))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAlgoServiceServer(srv, service)
	return srv
}
