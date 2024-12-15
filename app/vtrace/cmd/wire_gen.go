// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/data"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/server"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/fs/qiniuoss"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
)

// Injectors from wire.go:

func InitPlatform(ctx context.Context, c *conf.Bootstrap) (Platform, error) {
	etcd := &c.Etcd
	registryRegistry := registry.NewRegistry(ctx, etcd)
	dataData := data.NewData(ctx, registryRegistry)
	qiniuOSSConfig := c.OSS
	qiniuOSS := qiniuoss.NewQiNiuOSS(qiniuOSSConfig)
	vTraceCase := biz.NewVTraceCase(ctx, dataData, qiniuOSS)
	vTraceService := service.NewVTraceService(ctx, vTraceCase)
	grpcServer := server.NewGrpcServer(ctx, c, vTraceService)
	httpServer := server.NewHttpServer(ctx, c)
	platform := newApp(grpcServer, httpServer, registryRegistry)
	return platform, nil
}
