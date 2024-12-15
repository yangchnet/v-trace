// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/pkg/registry"

	"gitee.com/qciip-icp/v-trace/app/trans-job/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/trans-job/internal/job"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/cron"
	"github.com/google/wire"
)

func InitTransJob(ctx context.Context, c *conf.Bootstrap) (cron.JobInterface, error) {
	panic(wire.Build(
		wire.Bind(new(cron.JobInterface), new(*job.TransJob)),
		wire.Bind(new(cache.Cache), new(*cache.RedisStore)),
		wire.Bind(new(registry.Registrar), new(*registry.Registry)),
		wire.FieldsOf(new(*conf.Bootstrap), "Etcd"),
		job.ProviderSet,
		registry.RegistryProvider,
		NewConsumer,
		NewNodes,
	))
}
