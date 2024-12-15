package biz

import (
	"context"

	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	goodsV1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/data"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
	"gitee.com/qciip-icp/v-trace/pkg/tools/hashtools"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func (c *VTraceCase) CreateCirc(ctx context.Context, goodsId int32, circType string, from, to string, formValue []byte) (int32, string, error) {
	// 1. get cert
	sender := ctxtools.GetSenderFromCtx(ctx)
	caller, err := c.repo.GetCert(ctx, sender)
	if err != nil {
		return -1, "", err
	}

	var fromCert, toCert *transV1.Identity = &transV1.Identity{}, &transV1.Identity{}
	if from != "" {
		fromCert, err = c.repo.GetCert(ctx, from)
		if err != nil {
			return -1, "", err
		}
	}

	if to != "" {
		toCert, err = c.repo.GetCert(ctx, to)
		if err != nil {
			return -1, "", err
		}
	}

	transId, err := c.repo.TransId(ctx, goodsId)
	if err != nil {
		return -1, "", err
	}

	// 2. create circ
	circId, err := c.repo.CreateCirc(ctx, transId, circType, sender, from, to, formValue)
	if err != nil {
		return -1, "", err
	}

	// 3. do trans
	if err := c.repo.CallContract(
		ctx,
		&data.CallContractParams{
			MethodName: contract.CallAction[circType],
			Caller:     caller,
			From:       fromCert,
			To:         toCert,
			TokenId:    []int64{int64(goodsId)},
			TransId:    []string{transId},
			Digest:     hashtools.Sha256(formValue),
		}); err != nil {
		return -1, "", err
	}

	return circId, transId, nil
}

// 获取流转历史.
func (c *VTraceCase) GetGoodsAndCircs(ctx context.Context, goosId int32) (*goodsV1.Goods, []*circV1.CircRecord, error) {
	pbGoods, err := c.repo.GetGoodsByGoodsId(ctx, goosId)
	if err != nil {
		return nil, nil, err
	}

	pbCircs, err := c.repo.GetCircByGoodsId(ctx, goosId)
	if err != nil {
		return nil, nil, err
	}

	for _, circ := range pbCircs {
		pbTx, err := c.repo.GetTxByTransId(ctx, circ.GetTransId().GetValue())
		if err != nil {
			continue
		}
		circ.TxHash = pbtools.ToProtoString(pbTx.TxHash.GetValue())
	}

	return pbGoods, pbCircs, nil
}

// 批量流转
func (c *VTraceCase) BatchCirc(ctx context.Context, goodsIds []int32, circType, from, to string, formValue []byte) (map[int32]string, error) {
	// 1. get cert
	sender := ctxtools.GetSenderFromCtx(ctx)
	caller, err := c.repo.GetCert(ctx, sender)
	if err != nil {
		return nil, err
	}

	var fromCert, toCert *transV1.Identity = &transV1.Identity{}, &transV1.Identity{}
	if from != "" {
		fromCert, err = c.repo.GetCert(ctx, from)
		if err != nil {
			return nil, err
		}
	}

	if to != "" {
		toCert, err = c.repo.GetCert(ctx, to)
		if err != nil {
			return nil, err
		}
	}

	transIdMap, err := c.repo.BatchTransId(ctx, goodsIds)
	if err != nil {
		return nil, err
	}

	transIds := make([]string, 0)
	for _, v := range transIdMap {
		transIds = append(transIds, v)
	}

	// 2. create circ
	_, err = c.repo.BatchCirc(ctx, transIds, circType, sender, from, to, formValue)
	if err != nil {
		return nil, err
	}

	// 3. do trans
	for _, goodsId := range goodsIds {
		transId, ok := transIdMap[goodsId]
		if !ok {
			continue
		}
		if err := c.repo.CallContract(
			ctx,
			&data.CallContractParams{
				MethodName: contract.CallAction[circType],
				Caller:     caller,
				From:       fromCert,
				To:         toCert,
				TokenId:    []int64{int64(goodsId)},
				TransId:    []string{transId},
				Digest:     hashtools.Sha256(formValue),
			}); err != nil {
			return nil, err
		}
	}

	return transIdMap, nil
}
