package data

import (
	"context"
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"

	goodsV1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

// 创建商品类别.
func (r *Data) CreateGoodsClass(ctx context.Context, name string, des []byte, materials, orgId int32, tm string) (int32, error) {
	resp, err := r.Goods().CreateClass(withMetadata(ctx), &v1.CreateClassRequest{
		GoodsName: name,
		GoodsDes:  pbtools.ToProtoBytes(des),
		Material:  materials,
		OrgId:     orgId,
		Tm:        pbtools.ToProtoString(tm),
	})
	if err != nil {
		logger.Error(err)
		return -1, err
	}

	return int32(resp.GetGoodsId().GetValue()), nil
}

// 创建商品批次.
func (r *Data) CreateGoodsSerial(ctx context.Context, productTime time.Time, class_id int32) (int32, error) {
	resp, err := r.Goods().CreateSerial(withMetadata(ctx), &v1.CreateSerialRequest{
		ProductTime: pbtools.ToProtoTimestamp(productTime),
		ClassId:     class_id,
	})
	if err != nil {
		logger.Error(err)
		return -1, err
	}

	return int32(resp.GetSerialId().GetValue()), nil
}

// 批量创建商品.
func (r *Data) BatchCreateGoods(ctx context.Context, serial_id, sum int32) ([]int32, error) {
	resp, err := r.Goods().BatchCreateGoods(withMetadata(ctx), &goodsV1.BatchCreateGoodsRequest{
		SerialId: serial_id,
		Sum:      sum,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return pbtools.FromProtoInt32Slice(resp.GetIds()), nil
}

// 获取商品.
func (r *Data) GetGoodsByGoodsId(ctx context.Context, goods_id int32) (*goodsV1.Goods, error) {
	goods, err := r.Goods().GetGoods(ctxtools.WithMetadata(ctx), &goodsV1.GetGoodsRequest{
		GoodsId: goods_id,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return goods.Good, nil
}

// GetClass 获取类型
func (r *Data) GetClass(ctx context.Context, goods_id int32) (*goodsV1.Class, error) {
	class, err := r.Goods().GetClass(ctxtools.WithMetadata(ctx), &goodsV1.GetClassRequest{GoodsId: goods_id})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return class.Class, nil
}

// GetGoodsSerial 获取批次
func (r *Data) GetGoodsSerial(ctx context.Context, serial_id int32) (*goodsV1.Serial, error) {
	serial, err := r.Goods().GetSerial(ctxtools.WithMetadata(ctx), &goodsV1.GetSerialRequest{SerialId: serial_id})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return serial.Serial, nil
}

// 列出产品类型
func (r *Data) ListClass(ctx context.Context, offset, limit, orgId int32) ([]*goodsV1.Class, error) {
	classes, err := r.Goods().ListGoodsClass(ctxtools.WithMetadata(ctx), &goodsV1.ListGoodsClassRequest{
		Offset: offset,
		Limit:  limit,
		OrgId:  orgId,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return classes.GoodsClasses, nil
}

// 列出产品批次
func (r *Data) ListGoodsSerial(ctx context.Context, offset, limit, orgId int32) ([]*goodsV1.Serial, error) {
	serials, err := r.Goods().ListGoodsSerial(ctxtools.WithMetadata(ctx), &goodsV1.ListGoodsSerialRequest{
		Offset: offset,
		Limit:  limit,
		OrgId:  orgId,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return serials.Serial, nil
}

// 列出商品
func (r *Data) ListGoods(ctx context.Context, offset, limit, orgId int32) ([]*goodsV1.Goods, error) {
	goods, err := r.Goods().ListGoods(ctxtools.WithMetadata(ctx), &goodsV1.ListGoodsRequest{
		Offset: offset,
		Limit:  limit,
		OrgId:  orgId,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return goods.Goods, nil
}

// 更新商品种类
func (r *Data) UpdateClass(ctx context.Context, class *goodsV1.Class) error {
	if _, err := r.Goods().UpdateGoodsClass(ctxtools.WithMetadata(ctx), &goodsV1.UpdateGoodsClassRequest{Class: class}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// 更新产品批次
func (r *Data) UpdateGoodsSerial(ctx context.Context, serial *goodsV1.Serial) error {
	if _, err := r.Goods().UpdateGoodsSerial(ctxtools.WithMetadata(ctx), &goodsV1.UpdateGoodsSerialRequest{Serial: serial}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// 更新商品
func (r *Data) UpdateGoods(ctx context.Context, goods *goodsV1.Goods) error {
	if _, err := r.Goods().UpdateGoods(ctxtools.WithMetadata(ctx), &goodsV1.UpdateGoodsRequest{Goods: goods}); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// 获取所属企业
func (r *Data) GetOrgOfClass(ctx context.Context, classId int32) (int32, error) {
	resp, err := r.Goods().GetOrgOfX(ctx, &v1.GetOrgOfXRequest{
		X:  goodsV1.GetOrgOfXRequest_class,
		Id: classId,
	})
	if err != nil {
		logger.Error(err)

		return -1, err
	}

	return resp.GetOrgId().GetValue(), nil
}
func (r *Data) GetOrgOfSerial(ctx context.Context, serialId int32) (int32, error) {
	resp, err := r.Goods().GetOrgOfX(ctx, &v1.GetOrgOfXRequest{
		X:  goodsV1.GetOrgOfXRequest_serial,
		Id: serialId,
	})
	if err != nil {
		logger.Error(err)

		return -1, err
	}

	return resp.GetOrgId().GetValue(), nil
}
func (r *Data) GetOrgOfGoods(ctx context.Context, goodsId int32) (int32, error) {
	resp, err := r.Goods().GetOrgOfX(ctx, &v1.GetOrgOfXRequest{
		X:  goodsV1.GetOrgOfXRequest_goods,
		Id: goodsId,
	})
	if err != nil {
		logger.Error(err)

		return -1, err
	}

	return resp.GetOrgId().GetValue(), nil
}
