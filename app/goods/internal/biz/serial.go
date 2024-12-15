package biz

import (
	"context"
	"database/sql"
	"errors"
	"time"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/ctxtools"
)

// 创建产品批次.
func (c *GoodsCase) CreateSerial(ctx context.Context, productTime time.Time, classId int32) (*db.Serial, error) {
	creator := ctxtools.GetSenderFromCtx(ctx)
	var serialId int64
	if err := c.repo.ExecTx(ctx, func(q *db.Queries) error {
		_, err := q.GetClassById(ctx, classId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return v1.ErrorGoodsClassNotFound("不存在的产品种类ID: [%d]", classId)
			}

			return err
		}

		id, err := c.repo.CreateSerial(ctx, &db.CreateSerialParams{
			ProductTime: sql.NullTime{
				Time: productTime,

				Valid: true,
			},
			Creator: sql.NullString{
				String: creator,
				Valid:  true,
			},
			ClassID: sql.NullInt32{
				Int32: int32(classId),
				Valid: true,
			},
		})
		if err != nil {
			return v1.ErrorGoodsSerialCreateFailed("创建产品批次失败: %v", err)
		}

		serialId = id

		return nil
	}); err != nil {
		logger.Error(err)

		return nil, err
	}

	key := cache.GenKey("goods", "serial", "id", serialId)
	serialI, err := c.repo.CacheCreate(ctx, key, func(q db.Querier) (any, error) {
		serial, err := c.repo.GetSerialById(ctx, int32(serialId))
		if err != nil {
			logger.Error(err)

			return nil, v1.ErrorGoodsSerialNotFound("不存在的产品批次:%d", serialId)
		}

		return serial, nil
	})
	if err != nil {
		return nil, err
	}
	return serialI.(*db.Serial), nil
}

// 获取产品批次
func (c *GoodsCase) GetSerial(ctx context.Context, id int32) (*db.Serial, *db.Class, error) {
	key := cache.GenKey("goods", "serial", "id", id)
	serialI, err := c.repo.CacheGet(ctx, key, func(q db.Querier) (any, error) {
		serial, err := c.repo.GetSerialById(ctx, int32(id))
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsSerialNotFound("未找到产品批次: %d\n", id)
			}
			return nil, err
		}
		return serial, nil
	})
	if err != nil {
		return nil, nil, err
	}

	serial := serialI.(*db.Serial)

	serialKey := cache.GenKey("goods", "class", "id", serial.ClassID.Int32)
	classI, err := c.repo.CacheGet(ctx, serialKey, func(q db.Querier) (any, error) {
		class, err := c.repo.GetClassById(ctx, serial.ClassID.Int32)
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsClassNotFound("未找到产品批次: %d\n", id)
			}
			return nil, err
		}

		return class, nil
	})
	if err != nil {
		return nil, nil, err
	}

	return serial, classI.(*db.Class), nil
}

// 列出产品批次
func (c *GoodsCase) ListGoodsSerial(ctx context.Context, offset, limit, orgID int32) ([]*db.Serial, []*db.Class, error) {
	params, err := c.repo.ListGoodsSerial(ctx, &db.ListGoodsSerialParams{
		OrgID:  orgID,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		logger.Error(err)

		return nil, nil, err
	}

	serials := make([]*db.Serial, 0)
	classes := make([]*db.Class, 0)
	for _, param := range params {
		serials = append(serials, &db.Serial{
			ID:          param.ID,
			ProductTime: param.ProductTime,
			Status:      param.Status,
			CreatedAt:   param.CreatedAt,
			Creator:     param.Creator,
			ClassID:     param.ClassID,
		})
		classes = append(classes, &db.Class{
			ID:         param.ID_2,
			Name:       param.Name,
			Des:        param.Des,
			Status:     param.Status_2,
			CreatedAt:  param.CreatedAt_2,
			Creator:    param.Creator_2,
			MaterialID: param.MaterialID,
			OrgID:      param.OrgID,
			Tm:         param.Tm,
		})
	}

	return serials, classes, nil
}

// 更新产品批次
func (c *GoodsCase) UpdateGoodsSerial(ctx context.Context, serial *db.Serial) (*db.Serial, error) {
	key := cache.GenKey("goods", "serial", "id", int32(serial.ID))
	if err := c.repo.ExecTx(ctx, func(queries *db.Queries) error {
		return c.repo.CacheUpdate(ctx, key, func(q db.Querier) error {
			if err := queries.UpdateGoodSerial(ctx, &db.UpdateGoodSerialParams{
				ProductTime: sql.NullTime{
					Time:  serial.ProductTime.Time,
					Valid: true,
				},
			}); err != nil {
				logger.Error(err)
				if errors.Is(err, sql.ErrNoRows) {
					return v1.ErrorGoodsSerialNotFound("待更新产品批次未找到: %v\n", err)
				}
				return err
			}
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return serial, nil
}
