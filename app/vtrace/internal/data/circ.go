package data

import (
	"context"

	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func (r *Data) CreateCirc(ctx context.Context, transId string, circType string, operator, from, to string, fromInfo []byte) (int32, error) {
	resp, err := r.Circ().CreateCirc(ctx, &circV1.CreateCircRequest{
		TransId:  transId,
		CircType: v1.CircType(circV1.CircType_value[circType]),
		Operator: operator,
		From:     pbtools.ToProtoString(from),
		To:       to,
		FormInfo: pbtools.ToProtoBytes(fromInfo),
	})
	if err != nil {
		logger.Error(err)
		return -1, err
	}

	return int32(resp.GetCircId().GetValue()), nil
}

// 获取流转历史.
func (r *Data) GetCircByGoodsId(ctx context.Context, goods_id int32) ([]*circV1.CircRecord, error) {
	resp, err := r.Circ().GetCircByGoodsId(ctx, &v1.GetCircByGoodsIdRequest{
		GoodsId: int64(goods_id),
	})
	if err != nil {
		return nil, err
	}

	return resp.Records, nil
}

// 生成TransId
func (r *Data) TransId(ctx context.Context, goodsId int32) (string, error) {
	resp, err := r.Circ().TransId(ctx, &v1.TransIdRequest{
		GoodsId: goodsId,
	})

	if err != nil {
		return "", err
	}

	return resp.TransId, nil
}

func (r *Data) BatchTransId(ctx context.Context, goodsIds []int32) (map[int32]string, error) {
	resp, err := r.Circ().BatchTransId(ctx, &v1.BatchTransIdRequest{
		GoodsIds: goodsIds,
	})
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return resp.GetTransIds(), nil
}

func (r *Data) BatchCirc(ctx context.Context, transIds []string, circType string, operator, from, to string, fromInfo []byte) ([]int32, error) {
	resp, err := r.Circ().BatchCirc(ctx, &v1.BatchCircRequest{
		TransIds: transIds,
		CircType: v1.CircType(circV1.CircType_value[circType]),
		Operator: operator,
		From:     pbtools.ToProtoString(from),
		To:       pbtools.ToProtoString(to),
		FormInfo: pbtools.ToProtoBytes(fromInfo),
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return resp.GetCircIds(), nil
}
