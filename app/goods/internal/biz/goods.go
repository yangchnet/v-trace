package biz

import (
	"context"
	"database/sql"
	"errors"

	"gitee.com/qciip-icp/v-trace/pkg/cache"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
)

// BatchCreateGoods 批量创建产品.
func (c *GoodsCase) BatchCreateGoods(ctx context.Context, serialId int32, sum int32) ([]int32, error) {
	creator := ctxtools.GetSenderFromCtx(ctx)
	var ids []int32
	if err := c.repo.ExecTx(ctx, func(queries *db.Queries) error {
		for i := 0; i < int(sum); i++ {
			id, err := queries.CreateGoods(ctx, &db.CreateGoodsParams{
				Status: sql.NullString{
					String: "created",
					Valid:  true,
				},
				Creator: sql.NullString{
					String: creator,
					Valid:  true,
				},
				SerialID: sql.NullInt32{
					Int32: int32(serialId),
					Valid: true,
				},
			})
			if err != nil {
				logger.Error(err)

				return v1.ErrorGoodsCreateFailed("创建产品失败: %v", err)
			}

			ids = append(ids, int32(id))
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return ids, nil
}

// GetGoods 获取商品.
func (c *GoodsCase) GetGoods(ctx context.Context, goodsId int32) (*db.Good, *db.Serial, *db.Class, error) {
	keyGoods := cache.GenKey("goods", "goods", "id", goodsId)
	goodsI, err := c.repo.CacheGet(ctx, keyGoods, func(q db.Querier) (any, error) {
		goods, err := c.repo.GetGoodsById(ctx, int32(goodsId))
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsNotFound("未找到产品ID: %d", goodsId)
			}

			return nil, err
		}
		return goods, nil
	})
	if err != nil {
		return nil, nil, nil, err
	}
	keySerial := cache.GenKey("goods", "serial", "id", int32(goodsI.(*db.Good).SerialID.Int32))
	serialI, err := c.repo.CacheGet(ctx, keySerial, func(q db.Querier) (any, error) {
		serial, err := c.repo.GetSerialById(ctx, goodsI.(*db.Good).SerialID.Int32)
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsNotFound("未找到产品批次ID: %d", goodsId)
			}

			return nil, err
		}
		return serial, nil
	})
	if err != nil {
		return nil, nil, nil, err
	}
	keyClass := cache.GenKey("goods", "class", "id", int32(serialI.(*db.Serial).ClassID.Int32))
	classI, err := c.repo.CacheGet(ctx, keyClass, func(q db.Querier) (any, error) {
		class, err := c.repo.GetClassById(ctx, serialI.(*db.Serial).ClassID.Int32)
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsNotFound("未找到产品种类ID: %d", goodsId)
			}

			return nil, err
		}
		return class, nil
	})
	if err != nil {
		return nil, nil, nil, err
	}
	return goodsI.(*db.Good), serialI.(*db.Serial), classI.(*db.Class), nil
}

// ListGoods 列出商品
func (c *GoodsCase) ListGoods(ctx context.Context, offset, limit, orgId int32) ([]*db.Good, []*db.Serial, []*db.Class, error) {
	params, err := c.repo.ListGoods(ctx, &db.ListGoodsParams{
		OrgID:  orgId,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		logger.Error(err)

		return nil, nil, nil, err
	}

	goods := make([]*db.Good, 0)
	serials := make([]*db.Serial, 0)
	classes := make([]*db.Class, 0)
	for _, param := range params {
		goods = append(goods, &db.Good{
			ID:        param.ID,
			Status:    param.Status,
			CreatedAt: param.CreatedAt,
			Creator:   param.Creator,
			SerialID:  param.SerialID,
		})
		serials = append(serials, &db.Serial{
			ID:          param.ID_2,
			ProductTime: param.ProductTime,
			Status:      param.Status_2,
			CreatedAt:   param.CreatedAt_2,
			Creator:     param.Creator_2,
			ClassID:     param.ClassID,
		})
		classes = append(classes, &db.Class{
			ID:         param.ID_3,
			Name:       param.Name,
			Des:        param.Des,
			Status:     param.Status_3,
			CreatedAt:  param.CreatedAt_3,
			Creator:    param.Creator_3,
			MaterialID: param.MaterialID,
			OrgID:      param.OrgID,
			Tm:         param.Tm,
		})
	}

	return goods, serials, classes, nil
}

// UpdateGoods 更新产品
func (c *GoodsCase) UpdateGoods(ctx context.Context, goods *db.Good) (*db.Good, error) {
	key := cache.GenKey("goods", "goods", "id", goods.ID)
	if err := c.repo.ExecTx(ctx, func(queries *db.Queries) error {
		if err := c.repo.CacheUpdate(ctx, key, func(q db.Querier) error {
			if err := queries.UpdateGood(ctx, &db.UpdateGoodParams{
				Status: sql.NullString{
					String: goods.Status.String,
					Valid:  true,
				},
				Creator: sql.NullString{
					String: goods.Creator.String,
					Valid:  true,
				},
				SerialID: sql.NullInt32{
					Int32: goods.SerialID.Int32,
					Valid: true,
				},
			}); err != nil {
				logger.Error(err)
				return v1.ErrorGoodsNotFound("待更新产品未找到: %v\n", err)
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	keyGoods := cache.GenKey("goods", "goods", "id", goods.ID)
	goodsI, err := c.repo.CacheGet(ctx, keyGoods, func(q db.Querier) (any, error) {
		goods, err := c.repo.GetGoodsById(ctx, int32(goods.ID))
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsNotFound("未找到产品ID: %d", goods.ID)
			}

			return nil, err
		}
		return goods, nil
	})
	if err != nil {
		return nil, err
	}

	return goodsI.(*db.Good), nil
}
