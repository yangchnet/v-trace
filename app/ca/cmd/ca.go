package main

import (
	"context"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/tools/idtools"

	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models/db"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/services"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/utils"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
)

func newApp(grpcServer *grpc.Server, r registry.Registrar) *app.App {
	return app.New(
		app.WithServer(grpcServer),
		app.WithID(idtools.NewId()),
		app.WithName(constants.CA),
		app.WithVersion("v0.0.1"),
		app.WithRegistrar(r),
		app.WithRegistrarTimeout(time.Second*5),
	)
}

var allConfig *utils.AllConfig

func init() {
	allConfig = utils.SetConfig("app/ca/config/configs.yaml")
	db.GormInit()
	services.InitServer()
}

func main() {
	ctx := context.Background()

	app := InitApp(ctx, allConfig)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
