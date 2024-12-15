// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/biz"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/server"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/app"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
)

// Injectors from wire.go:

func InitApp(ctx context.Context, c *conf.Bootstrap) (*app.App, error) {
	dataConfig := &c.Data
	dataData := data.NewData(ctx, dataConfig)
	goodsCase := biz.NewGoodsCase(dataData)
	goodsService := service.NewGoodsService(ctx, goodsCase)
	grpcServer := server.NewGrpcServer(ctx, c, goodsService)
	etcd := &c.Etcd
	registryRegistry := registry.NewRegistry(ctx, etcd)
	appApp := newApp(grpcServer, registryRegistry)
	return appApp, nil
}