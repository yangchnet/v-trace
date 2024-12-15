package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/server"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/tools/idtools"
	"golang.org/x/sync/errgroup"
	"time"

	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/conf"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
)

type Platform map[string]*app.App

func newApp(grpcServer *grpc.Server, httpServer *server.HttpServer, r registry.Registrar) Platform {
	grpc := app.New(
		app.WithServer(grpcServer),
		app.WithID(idtools.NewId()),
		app.WithName(constants.VTrace),
		app.WithVersion("v0.0.1"),
		app.WithRegistrar(r),
		app.WithRegistrarTimeout(time.Second*5),
	)

	http := app.New(
		app.WithID(idtools.NewId()),
		app.WithName("vtrace-http"),
		app.WithServer(httpServer),
		app.WithVersion("v0.0.1"),
	)

	pf := make(Platform)
	pf["grpc"] = grpc
	pf["http"] = http

	return pf
}

func main() {
	ctx := context.Background()

	c := conf.LoadConfig("app/vtrace/config")

	pf, err := InitPlatform(ctx, c)
	if err != nil {
		panic(err)
	}

	eg := errgroup.Group{}

	eg.Go(func() error {
		return pf["grpc"].Run()
	})

	time.Sleep(time.Second * 5) // wait for grpc server registered

	eg.Go(func() error {
		return pf["http"].Run()
	})

	eg.Wait()
}
