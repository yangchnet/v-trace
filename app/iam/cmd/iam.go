package main

import (
	"context"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/tools/idtools"

	"gitee.com/qciip-icp/v-trace/app/iam/internal/conf"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
)

func newApp(grpcServer *grpc.Server, r registry.Registrar) *app.App {
	return app.New(
		app.WithServer(grpcServer),
		app.WithID(idtools.NewId()),
		app.WithName("iam"),
		app.WithVersion("v0.0.1"),
		app.WithRegistrar(r),
		app.WithRegistrarTimeout(time.Second*5),
	)
}

func main() {
	// time.Sleep(10 * time.Second)
	ctx := context.Background()

	c := conf.LoadConfig("app/iam/config")

	app, err := InitApp(ctx, c)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
