//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/data"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/server"
	"gitee.com/qciip-icp/v-trace/app/iam/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/token"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, c *conf.Bootstrap) (*app.App, error) {
	panic(wire.Build(
		wire.Bind(new(biz.IamRepo), new(*data.Data)),
		wire.Bind(new(service.IamCaseInterface), new(*biz.IamCase)),
		wire.Bind(new(v1.IamServiceServer), new(*service.IamService)),
		wire.FieldsOf(new(*conf.Bootstrap), "Data"),
		wire.FieldsOf(new(*conf.Bootstrap), "Token"),
		newApp,
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		registry.RegistryProvider,
		server.ProviderSet,
		service.ProviderSet,
		biz.ProviderSet,
		token.ProviderSet,
		data.ProviderSet,
	))
}
