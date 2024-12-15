package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/conf"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/tools/idtools"
	"time"
)

func main() {
	ctx := context.Background()

	c := conf.LoadConfig("app/algo/config")

	app, err := InitApp(ctx, c)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(grpcServer *grpc.Server, r registry.Registrar) *app.App {
	return app.New(
		app.WithServer(grpcServer),
		app.WithID(idtools.NewId()),
		app.WithName(constants.Algo),
		app.WithVersion("v0.0.1"),
		app.WithRegistrar(r),
		app.WithRegistrarTimeout(time.Second*5),
	)
}
