//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/pkg/fs/qiniuoss"

	"gitee.com/qciip-icp/v-trace/pkg/registry"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/data"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/server"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/fs"
	"github.com/google/wire"
)

func InitPlatform(ctx context.Context, c *conf.Bootstrap) (Platform, error) {
	panic(wire.Build(
		wire.Bind(new(biz.VTraceRepo), new(*data.Data)),
		wire.Bind(new(service.VTraceCaseInterface), new(*biz.VTraceCase)),
		wire.Bind(new(v1.VTraceInterfaceServer), new(*service.VTraceService)),
		wire.Bind(new(fs.Interface), new(*qiniuoss.QiniuOSS)),
		newApp,
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		wire.FieldsOf(new(*conf.Bootstrap), "OSS"),
		registry.RegistryProvider,
		server.ProviderSet,
		server.HttpProvider,
		service.ProviderSet,
		biz.ProviderSet,
		qiniuoss.OssProvider,
		data.ProviderSet,
	))
}
