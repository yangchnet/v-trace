package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

// 商品流转.
func (s *VTraceService) CreateCirc(ctx context.Context, req *v1.CreateCircRequest) (*v1.CreateCircResponse, error) {
	goodsId := req.GetGoodsId()
	circType := req.GetCircType().String()
	from := req.GetFrom().GetValue()
	to := req.GetTo()
	formValue := req.GetFormValue().GetValue()

	circId, transId, err := s.cas.CreateCirc(ctx, goodsId, circType, from, to, formValue)
	if err != nil {
		return nil, err
	}

	return &v1.CreateCircResponse{
		CircId:  pbtools.ToProtoInt64(circId),
		TransId: pbtools.ToProtoString(transId),
	}, nil
}

// 获取商品流转历史.
func (s *VTraceService) GetCirc(ctx context.Context, req *v1.GetCircRequest) (*v1.GetCircResponse, error) {
	goods, records, err := s.cas.GetGoodsAndCircs(ctx, req.GetGoodsId())
	if err != nil {
		return nil, err
	}

	resp := &v1.GetCircResponse{
		GoodsInfo:   goods,
		CircRecords: records,
	}

	return resp, nil
}

// 产品批量流传
func (s *VTraceService) BatchCirc(ctx context.Context, req *v1.BatchCircRequest) (*v1.BatchCircResponse, error) {
	m, err := s.cas.BatchCirc(ctx, req.GetGoodsIds(), req.GetCircType().String(), req.GetFrom().GetValue(), req.GetTo().GetValue(), req.GetFormInfo().GetValue())
	if err != nil {
		return nil, err
	}

	return &v1.BatchCircResponse{
		GoodsId2TransId: m,
	}, nil
}
