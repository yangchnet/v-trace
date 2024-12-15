//go:build wireinject
// +build wireinject

package main

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/data"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/server"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, c *conf.Bootstrap) (*app.App, error) {
	panic(wire.Build(
		wire.Bind(new(biz.TransRepo), new(*data.Data)),
		wire.Bind(new(service.TransCaseInterface), new(*biz.TransCase)),
		wire.Bind(new(v1.TransServiceServer), new(*service.TransService)),
		wire.FieldsOf(new(*conf.Bootstrap), "Data"),
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		registry.RegistryProvider,
		newApp,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	))
}
