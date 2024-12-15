package server

import (
	"context"
	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/conf"
	localGrpc "gitee.com/qciip-icp/v-trace/pkg/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGrpcServer)

func NewGrpcServer(ctx context.Context, c *conf.Bootstrap, service v1.IamServiceServer) *localGrpc.Server {
	opts := make([]localGrpc.ServerOption, 0)
	if c.Grpc.Network != "" {
		opts = append(opts, localGrpc.WithNetwork(c.Grpc.Network))
	}
	if c.Grpc.Address != "" {
		opts = append(opts, localGrpc.WithAddress(c.Grpc.Address))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, localGrpc.WithTimeout(*c.Grpc.Timeout))
	}

	srv := localGrpc.NewServer(opts...)
	v1.RegisterIamServiceServer(srv, service)

	return srv
}
