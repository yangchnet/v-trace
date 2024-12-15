package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

func (s *GoodsService) CreateSerial(ctx context.Context, req *v1.CreateSerialRequest) (*v1.CreateSerialResponse, error) {
	productTime := req.GetProductTime().AsTime()
	classId := req.GetClassId()

	serial, err := s.cas.CreateSerial(ctx, productTime, classId)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.CreateSerialResponse{
		SerialId: pbtools.ToProtoInt32(serial.ID),
	}, nil
}

// 获取批次
func (s *GoodsService) GetSerial(ctx context.Context, req *v1.GetSerialRequest) (*v1.GetSerialResponse, error) {
	serial, class, err := s.cas.GetSerial(ctx, req.SerialId)
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.GetSerialResponse{Serial: Serial2Proto(serial, class)}, nil
}

// 列出产品批次
func (s *GoodsService) ListGoodsSerial(ctx context.Context, req *v1.ListGoodsSerialRequest) (*v1.ListGoodsSerialResponse, error) {
	serials, classes, err := s.cas.ListGoodsSerial(ctx, req.Offset, req.Limit, req.GetOrgId())
	if err != nil {
		return nil, verr.Error(s, err)
	}
	res := make([]*v1.Serial, 0)
	for i, serial := range serials {
		res = append(res, Serial2Proto(serial, classes[i]))
	}
	return &v1.ListGoodsSerialResponse{Serial: res}, nil
}

// 更新产品批次
func (s *GoodsService) UpdateGoodsSerial(ctx context.Context, req *v1.UpdateGoodsSerialRequest) (*v1.UpdateGoodsSerialResponse, error) {
	_, err := s.cas.UpdateGoodsSerial(ctx, Proto2Serial(req.GetSerial()))
	if err != nil {
		return nil, verr.Error(s, err)
	}
	return &v1.UpdateGoodsSerialResponse{Ok: true}, nil
}
