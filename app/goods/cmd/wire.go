//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/server"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, c *conf.Bootstrap) (*app.App, error) {
	panic(wire.Build(
		wire.Bind(new(biz.GoodsRepo), new(*data.Data)),
		wire.Bind(new(service.GoodsCaseInterface), new(*biz.GoodsCase)),
		wire.Bind(new(v1.GoodsServiceServer), new(*service.GoodsService)),
		wire.FieldsOf(new(*conf.Bootstrap), "Data"),
		newApp,
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		registry.RegistryProvider,
		server.ProviderSet,
		service.ProviderSet,
		biz.ProviderSet,
		data.ProviderSet,
	))
}
