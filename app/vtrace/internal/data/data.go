package data

import (
	"context"
	goodsV1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	localGrpc "gitee.com/qciip-icp/v-trace/pkg/grpc"

	algoV1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	caV1 "gitee.com/qciip-icp/v-trace/api/ca/v1"
	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	iamV1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"github.com/google/wire"
	"google.golang.org/grpc/metadata"
)

var ProviderSet = wire.NewSet(NewData)

// Data handle other services' clients.
type Data struct {
	getters map[string]localGrpc.ConnGetInterface

	r registry.Registrar
}

// NewData build a new Data.
func NewData(ctx context.Context, r registry.Registrar) *Data {
	getters := make(map[string]localGrpc.ConnGetInterface)

	iamWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Iam)
	if err != nil {
		panic(err)
	}
	getters[constants.Iam] = localGrpc.NewConnGetter(ctx, iamWatcher, constants.Iam)

	caWatcher, err := r.(*registry.Registry).Watch(ctx, constants.CA)
	if err != nil {
		panic(err)
	}
	getters[constants.CA] = localGrpc.NewConnGetter(ctx, caWatcher, constants.CA)

	transWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Trans)
	if err != nil {
		panic(err)
	}
	getters[constants.Trans] = localGrpc.NewConnGetter(ctx, transWatcher, constants.Trans)

	goodsWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Goods)
	if err != nil {
		panic(err)
	}
	getters[constants.Goods] = localGrpc.NewConnGetter(ctx, goodsWatcher, constants.Goods)

	circWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Circ)
	if err != nil {
		panic(err)
	}
	getters[constants.Circ] = localGrpc.NewConnGetter(ctx, circWatcher, constants.Circ)

	algoWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Algo)
	if err != nil {
		panic(err)
	}
	getters[constants.Algo] = localGrpc.NewConnGetter(ctx, algoWatcher, constants.Algo)

	return &Data{
		getters: getters,
		r:       r,
	}
}

func withMetadata(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Info("have no metadata")
	}

	return metadata.NewOutgoingContext(ctx, md)
}

func (d *Data) Iam() iamV1.IamServiceClient {
	conn := d.getters[constants.Iam].Get()

	return iamV1.NewIamServiceClient(conn)
}

func (d *Data) Ca() caV1.CAServiceClient {
	conn := d.getters[constants.CA].Get()

	return caV1.NewCAServiceClient(conn)
}

func (d *Data) Trans() transV1.TransServiceClient {
	conn := d.getters[constants.Trans].Get()

	return transV1.NewTransServiceClient(conn)
}

func (d *Data) Circ() circV1.CircServiceClient {
	conn := d.getters[constants.Circ].Get()

	return circV1.NewCircServiceClient(conn)
}

func (d *Data) Algo() algoV1.AlgoServiceClient {
	conn := d.getters[constants.Algo].Get()

	return algoV1.NewAlgoServiceClient(conn)
}

func (d *Data) Goods() goodsV1.GoodsServiceClient {
	conn := d.getters[constants.Goods].Get()

	return goodsV1.NewGoodsServiceClient(conn)
}
