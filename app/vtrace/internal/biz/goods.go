package biz

import (
	"context"
	"time"

	goodsV1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"

	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/data"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
	"gitee.com/qciip-icp/v-trace/pkg/tools/t2"
)

// 创建商品类别.
func (uc *VTraceCase) CreateGoodsClass(ctx context.Context, name string, des []byte, material int32, tm string) (int32, error) {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return -1, err
	}

	return uc.repo.CreateGoodsClass(ctx, name, des, material, int32(org.GetId().GetValue()), tm)
}

// 获取种类
func (uc *VTraceCase) GetClass(ctx context.Context, classId int32) (*goodsV1.Class, error) {
	return uc.repo.GetClass(ctx, classId)
}

// 列出商品种类
func (uc *VTraceCase) ListClass(ctx context.Context, offset, limit int32) ([]*goodsV1.Class, error) {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return uc.repo.ListClass(ctx, offset, limit, int32(org.GetId().GetValue()))
}

// 更新商品种类
func (uc *VTraceCase) UpdateClass(ctx context.Context, class *goodsV1.Class) error {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return err
	}

	orgId, err := uc.repo.GetOrgOfClass(ctx, int32(class.GetID().GetValue()))
	if err != nil {
		return err
	}

	if org.GetId().GetValue() != int64(orgId) {
		return v1.ErrorPermissionDeny("无权限进行该操作")
	}
	return uc.repo.UpdateClass(ctx, class)
}

// 创建商品批次.
func (uc *VTraceCase) CreateGoodsSerial(ctx context.Context, productTime time.Time, class_id int32) (int32, error) {
	return uc.repo.CreateGoodsSerial(ctx, productTime, class_id)
}

// 获得商品批次
func (uc *VTraceCase) GetGoodsSerial(ctx context.Context, serial_id int32) (*goodsV1.Serial, error) {
	return uc.repo.GetGoodsSerial(ctx, serial_id)
}

// 列出产品批次
func (uc *VTraceCase) ListGoodsSerial(ctx context.Context, offset, limit int32) ([]*goodsV1.Serial, error) {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return uc.repo.ListGoodsSerial(ctx, offset, limit, int32(org.GetId().GetValue()))
}

// 更新产品批次
func (uc *VTraceCase) UpdateGoodsSerial(ctx context.Context, serial *goodsV1.Serial) error {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return err
	}

	orgId, err := uc.repo.GetOrgOfSerial(ctx, int32(serial.GetID().GetValue()))
	if err != nil {
		return err
	}

	if org.GetId().GetValue() != int64(orgId) {
		return v1.ErrorPermissionDeny("无权限进行该操作")
	}
	return uc.repo.UpdateGoodsSerial(ctx, serial)
}

// 批量创建商品.
func (uc *VTraceCase) BatchCreateGoods(ctx context.Context, serial_id int32, sum int32) ([]int32, error) {
	sender := ctxtools.GetSenderFromCtx(ctx)

	// 1. create goods
	ids, err := uc.repo.BatchCreateGoods(ctx, serial_id, sum)
	if err != nil {
		return nil, err
	}

	// 2. get cert
	call, err := uc.repo.GetCert(ctx, "admin")
	if err != nil {
		return nil, err
	}

	to, err := uc.repo.GetCert(ctx, sender)
	if err != nil {
		return nil, err
	}

	// 3. get transId
	transIdMap, err := uc.repo.BatchTransId(ctx, ids)
	if err != nil {
		return nil, err
	}

	transIds := make([]string, 0)
	for _, v := range transIdMap {
		transIds = append(transIds, v)
	}

	_, err = uc.repo.BatchCirc(ctx, transIds, circV1.CircType_produce.String(), sender, "", sender, nil)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.CallContract(
		ctx,
		&data.CallContractParams{
			MethodName: contract.MethodBatchMint,
			Caller:     call,
			To:         to,
			TokenId:    t2.Int32ArrTo64(ids...),
			TransId:    transIds,
			Digest:     "",
		},
	); err != nil {
		return nil, err
	}

	return ids, nil
}

// 获得产品
func (uc *VTraceCase) GetGoods(ctx context.Context, goods_id int32) (*goodsV1.Goods, error) {
	return uc.repo.GetGoodsByGoodsId(ctx, goods_id)
}

// 列出商品
func (uc *VTraceCase) ListGoods(ctx context.Context, offset, limit int32) ([]*goodsV1.Goods, error) {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return uc.repo.ListGoods(ctx, offset, limit, int32(org.GetId().GetValue()))
}

// 更新商品
func (uc *VTraceCase) UpdateGoods(ctx context.Context, goods *goodsV1.Goods) error {
	org, err := uc.repo.GetOrgOfUser(ctx, ctxtools.GetSenderFromCtx(ctx))
	if err != nil {
		return err
	}

	orgId, err := uc.repo.GetOrgOfGoods(ctx, int32(goods.GetID().GetValue()))
	if err != nil {
		return err
	}

	if org.GetId().GetValue() != int64(orgId) {
		return v1.ErrorPermissionDeny("无权限进行该操作")
	}
	return uc.repo.UpdateGoods(ctx, goods)
}
