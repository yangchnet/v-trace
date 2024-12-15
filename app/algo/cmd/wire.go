//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/data"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/server"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/service"
	"gitee.com/qciip-icp/v-trace/app/pkg/algo"
	"gitee.com/qciip-icp/v-trace/pkg/registry"

	"gitee.com/qciip-icp/v-trace/pkg/app"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, c *conf.Bootstrap) (*app.App, error) {
	panic(wire.Build(
		wire.Bind(new(biz.AlgoRepo), new(*data.Data)),
		wire.Bind(new(service.AlgoCaseInterface), new(*biz.AlgoCase)),
		wire.Bind(new(v1.AlgoServiceServer), new(*service.AlgoService)),
		wire.FieldsOf(new(*conf.Bootstrap), "Algo"),
		wire.FieldsOf(new(*conf.Bootstrap), "Data"),
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		registry.RegistryProvider,
		server.ProviderSet,
		algo.AlgoProviderSet,
		service.ProviderSet,
		biz.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
