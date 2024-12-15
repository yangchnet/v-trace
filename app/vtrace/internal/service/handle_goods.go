package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

// 批量创建产品
func (s *VTraceService) BatchCreateGoods(ctx context.Context, req *v1.BatchCreateGoodsRequest) (*v1.BatchCreateGoodsResponse, error) {
	ids, err := s.cas.BatchCreateGoods(ctx, req.GetSerialId(), req.GetSum())
	if err != nil {
		return &v1.BatchCreateGoodsResponse{}, err
	}

	return &v1.BatchCreateGoodsResponse{
		Ids:       pbtools.ToProtoInt64List(ids),
		Successes: int32(len(ids)),
	}, nil
}

// 列出产品
func (s *VTraceService) ListGoods(ctx context.Context, req *v1.ListGoodsRequest) (*v1.ListGoodsResponse, error) {
	goods, err := s.cas.ListGoods(ctx, req.Offset, req.Limit)
	if err != nil {
		return &v1.ListGoodsResponse{}, err
	}
	return &v1.ListGoodsResponse{Goods: goods}, nil
}

// 更新产品
func (s *VTraceService) UpdateGoods(ctx context.Context, req *v1.UpdateGoodsRequest) (*v1.UpdateGoodsResponse, error) {
	err := s.cas.UpdateGoods(ctx, req.Goods)
	if err != nil {
		return &v1.UpdateGoodsResponse{}, err
	}
	return &v1.UpdateGoodsResponse{Ok: true}, nil
}

// 创建商品种类.
func (s *VTraceService) CreateGoodsClass(ctx context.Context, req *v1.CreateGoodsClassRequest) (*v1.CreateGoodsClassResponse, error) {
	class_id, err := s.cas.CreateGoodsClass(
		ctx,
		req.GetGoodsName(),
		req.GetGoodsDes().GetValue(),
		req.GetMaterialId(),
		req.GetTm().GetValue(),
	)
	if err != nil {
		logger.Error(err)
		return &v1.CreateGoodsClassResponse{
			ClassId: nil,
		}, err
	}

	return &v1.CreateGoodsClassResponse{
		ClassId: pbtools.ToProtoInt64(class_id),
	}, nil
}

// 列出产品类型
func (s *VTraceService) ListClass(ctx context.Context, req *v1.ListClassRequest) (*v1.ListClassResponse, error) {
	class, err := s.cas.ListClass(ctx, req.Offset, req.Limit)
	if err != nil {
		return &v1.ListClassResponse{}, err
	}
	return &v1.ListClassResponse{GoodsClasses: class}, nil
}

// 更新产品类型信息
func (s *VTraceService) UpdateClass(ctx context.Context, req *v1.UpdateClassRequest) (*v1.UpdateClassResponse, error) {
	err := s.cas.UpdateClass(ctx, req.Class)
	if err != nil {
		return &v1.UpdateClassResponse{}, err
	}
	return &v1.UpdateClassResponse{Ok: true}, nil
}

// 创建商品批次
func (s *VTraceService) CreateGoodsSerial(ctx context.Context, req *v1.CreateGoodsSerialRequest) (*v1.CreateGoodsSerialResponse, error) {
	productTime := req.GetProductTime().AsTime()
	serial_id, err := s.cas.CreateGoodsSerial(ctx, productTime, req.GetClassId())
	if err != nil {
		return &v1.CreateGoodsSerialResponse{
			SerialId: nil,
		}, err
	}

	return &v1.CreateGoodsSerialResponse{
		SerialId: pbtools.ToProtoInt64(serial_id),
	}, nil
}

// 列出产品批次
func (s *VTraceService) ListGoodsSerial(ctx context.Context, req *v1.ListGoodsSerialRequest) (*v1.ListGoodsSerialResponse, error) {
	serials, err := s.cas.ListGoodsSerial(ctx, req.Offset, req.Limit)
	if err != nil {
		return &v1.ListGoodsSerialResponse{}, err
	}
	return &v1.ListGoodsSerialResponse{Serial: serials}, nil
}

// 更新产品批次
func (s *VTraceService) UpdateGoodsSerial(ctx context.Context, req *v1.UpdateGoodsSerialRequest) (*v1.UpdateGoodsSerialResponse, error) {
	err := s.cas.UpdateGoodsSerial(ctx, req.Serial)
	if err != nil {
		return &v1.UpdateGoodsSerialResponse{}, err
	}
	return &v1.UpdateGoodsSerialResponse{Ok: true}, nil
}
