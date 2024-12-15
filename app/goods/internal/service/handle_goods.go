package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

// BatchCreateGoods 批量创建产品.
func (s *GoodsService) BatchCreateGoods(ctx context.Context, req *v1.BatchCreateGoodsRequest) (*v1.BatchCreateGoodsResponse, error) {
	serialId := req.GetSerialId()
	sum := req.GetSum()

	ids, err := s.cas.BatchCreateGoods(ctx, serialId, sum)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	if len(ids) == int(sum) {
		return &v1.BatchCreateGoodsResponse{
			Ids: pbtools.ToProtoInt32List(ids),
		}, nil
	}

	return &v1.BatchCreateGoodsResponse{
		Ids: pbtools.ToProtoInt32List(ids),
	}, nil
}

// 获得产品
func (s *GoodsService) GetGoods(ctx context.Context, req *v1.GetGoodsRequest) (*v1.GetGoodsResponse, error) {
	goods, serial, class, err := s.cas.GetGoods(ctx, req.GetGoodsId())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.GetGoodsResponse{Good: Goods2Proto(goods, serial, class)}, nil
}

// 列出产品
func (s *GoodsService) ListGoods(ctx context.Context, req *v1.ListGoodsRequest) (*v1.ListGoodsResponse, error) {
	goods, serial, classes, err := s.cas.ListGoods(ctx, req.Offset, req.Limit, req.GetOrgId())
	if err != nil {
		return nil, verr.Error(s, err)
	}
	res := make([]*v1.Goods, 0)
	for i, good := range goods {
		res = append(res, Goods2Proto(good, serial[i], classes[i]))
	}
	return &v1.ListGoodsResponse{Goods: res}, nil
}

// 更新产品
func (s *GoodsService) UpdateGoods(ctx context.Context, req *v1.UpdateGoodsRequest) (*v1.UpdateGoodsResponse, error) {
	goods := req.GetGoods()
	_, err := s.cas.UpdateGoods(ctx, Proto2Goods(goods))
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.UpdateGoodsResponse{Ok: true}, nil
}
