//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/utils"

	"gitee.com/qciip-icp/v-trace/app/ca/internal/server"
	"gitee.com/qciip-icp/v-trace/pkg/registry"

	"gitee.com/qciip-icp/v-trace/app/ca/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"github.com/google/wire"
)

func InitApp(ctx context.Context, c *utils.AllConfig) *app.App {
	panic(wire.Build(
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*utils.AllConfig), "Etcd"),
		registry.RegistryProvider,
		newApp,
		server.ProviderSet,
		service.ProviderSet,
	))
}
