package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

var _ v1.CircServiceServer = (*CircService)(nil)

// CreateCirc 创建一个流转记录.
func (s *CircService) CreateCirc(ctx context.Context, req *v1.CreateCircRequest) (*v1.CreateCircResponse, error) {
	record, err := s.cas.CreateRecord(
		ctx,
		req.GetTransId(),
		req.GetCircType().String(),
		req.GetOperator(),
		req.GetFrom().GetValue(),
		req.GetTo(),
		req.GetFormInfo().GetValue(),
	)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.CreateCircResponse{
		CircId: pbtools.ToProtoInt64(record.ID),
	}, nil
}

// BatchCirc 批量流转.
func (s *CircService) BatchCirc(ctx context.Context, req *v1.BatchCircRequest) (*v1.BatchCircResponse, error) {
	ids, err := s.cas.BatchRecord(
		ctx,
		req.GetTransIds(),
		req.GetOperator(),
		req.GetFrom().GetValue(),
		req.GetTo().GetValue(),
		req.GetCircType().String(),
		req.GetFormInfo().GetValue(),
	)
	if err != nil {
		return nil, verr.Error(s, err)
	}

	ids32 := make([]int32, len(ids))
	for i, id := range ids {
		ids32[i] = int32(id)
	}

	return &v1.BatchCircResponse{
		CircIds: ids32,
	}, nil
}

// GetCirc 获取一个流转记录.
func (s *CircService) GetCirc(ctx context.Context, req *v1.GetCircRequest) (*v1.CircRecord, error) {
	record, err := s.cas.GetRecordByID(ctx, int(req.GetCircId()))
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return Circ2Proto(record), nil
}

func (s *CircService) GetCircByTransId(ctx context.Context, req *v1.GetCircByTransIdRequest) (*v1.CircRecord, error) {
	record, err := s.cas.GetRecordByTransId(ctx, req.GetTransId())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return Circ2Proto(record), nil
}

// 根据商品id获取流转记录.
func (s *CircService) GetCircByGoodsId(ctx context.Context, req *v1.GetCircByGoodsIdRequest) (*v1.GetCircByGoodsIdResponse, error) {
	records, err := s.cas.GetRecordByGoodsId(ctx, int(req.GetGoodsId()))
	if err != nil {
		return nil, verr.Error(s, err)
	}

	pbRecords := make([]*v1.CircRecord, 0)
	for _, record := range records {
		pbRecords = append(pbRecords, Circ2Proto(record))
	}

	return &v1.GetCircByGoodsIdResponse{
		Records: pbRecords,
	}, nil
}

// 为商品生成transId
func (s *CircService) TransId(ctx context.Context, req *v1.TransIdRequest) (*v1.TransIdResponse, error) {
	transId, err := s.cas.TransId(ctx, req.GetGoodsId())
	if err != nil {
		return nil, err
	}

	return &v1.TransIdResponse{
		TransId: transId,
	}, nil
}

// 批量为商品生成transId
func (s *CircService) BatchTransId(ctx context.Context, req *v1.BatchTransIdRequest) (*v1.BatchTransIdResponse, error) {
	transIds, err := s.cas.BatchTransId(ctx, req.GetGoodsIds())
	if err != nil {
		return nil, err
	}

	return &v1.BatchTransIdResponse{
		TransIds: transIds,
	}, nil
}

// 更新流转记录状态
func (s *CircService) UpdateCircStatus(ctx context.Context, req *v1.UpdateCircStatusRequest) (*v1.UpdateCircStatusResponse, error) {
	if err := s.cas.UpdateCircStatus(ctx, req.GetTransId(), req.GetStatus().String()); err != nil {
		return nil, err
	}

	return &v1.UpdateCircStatusResponse{}, nil
}
