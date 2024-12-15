package biz

import (
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGoodsCase)

type GoodsCase struct {
	repo GoodsRepo
}

type GoodsRepo interface {
	data.Interface
}

func NewGoodsCase(goodsRepo GoodsRepo) *GoodsCase {
	return &GoodsCase{
		repo: goodsRepo,
	}
}
