package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGoodsService)

type GoodsService struct {
	v1.UnimplementedGoodsServiceServer

	cas GoodsCaseInterface
}

var _ v1.GoodsServiceServer = (*GoodsService)(nil)

func (s *GoodsService) mustEmbedUnimplementedGoodsServiceServer() {
	panic("not implemented") // 不需实现
}

func NewGoodsService(ctx context.Context, cas GoodsCaseInterface) *GoodsService {
	return &GoodsService{cas: cas}
}

func (s *GoodsService) GetDomain() string {
	return "goods-service"
}
