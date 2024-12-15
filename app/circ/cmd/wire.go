//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/data"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/server"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, c *conf.Bootstrap) (*app.App, error) {
	panic(wire.Build(
		wire.Bind(new(biz.CircRepo), new(*data.Data)),
		wire.Bind(new(service.CircCaseInterface), new(*biz.CircCase)),
		wire.Bind(new(v1.CircServiceServer), new(*service.CircService)),
		wire.FieldsOf(new(*conf.Bootstrap), "Data"),
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		registry.RegistryProvider,
		server.ProviderSet,
		service.ProviderSet,
		biz.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
