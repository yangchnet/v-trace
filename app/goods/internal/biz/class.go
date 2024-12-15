package biz

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
)

// 创建产品种类.
func (c *GoodsCase) CreateClass(ctx context.Context, name, creator string, des []byte, materialId, orgId int32, tm string) (*db.Class, error) {
	id, err := c.repo.CreateClass(ctx, &db.CreateClassParams{
		Name: sql.NullString{
			String: name,
			Valid:  true,
		},
		Des: des,
		Status: sql.NullString{
			String: "created",
			Valid:  true,
		},
		Creator: sql.NullString{
			String: creator,
			Valid:  true,
		},
		MaterialID: int32(materialId),
		OrgID:      int32(orgId),
		Tm: sql.NullString{
			String: tm,
			Valid:  true,
		},
	})
	if err != nil {
		logger.Error(err)
		if errors.Is(err, sql.ErrTxDone) {
			return nil, v1.ErrorGoodsClassCreateFailed("产品种类已存在: %v", err)
		}
		return nil, err
	}

	key := cache.GenKey("goods", "class", "id", id)
	classI, err := c.repo.CacheGet(ctx, key, func(q db.Querier) (any, error) {
		class, err := c.repo.GetClassById(ctx, int32(id))
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsClassNotFound("不存在产品类别ID:%d", id)
			}
			return nil, err
		}
		return class, nil
	})
	if err != nil {
		return nil, err
	}
	return classI.(*db.Class), nil
}

// GetGoodClassById 获取产品种类
func (c *GoodsCase) GetClass(ctx context.Context, id int32) (*db.Class, error) {
	key := cache.GenKey("goods", "class", "id", id)
	classI, err := c.repo.CacheGet(ctx, key, func(q db.Querier) (any, error) {
		class, err := c.repo.GetClassById(ctx, int32(id))
		if err != nil {
			logger.Error(err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, v1.ErrorGoodsClassNotFound("未找到产品种类: %d", id)
			}

			return nil, err
		}
		return class, nil
	})
	if err != nil {
		return nil, err
	}
	return classI.(*db.Class), nil
}

// ListGoodClass 列出产品类型
func (c *GoodsCase) ListGoodsClass(ctx context.Context, offset, limit, orgId int32) ([]*db.Class, error) {
	return c.repo.ListGoodClass(ctx, &db.ListGoodClassParams{
		OrgID:  int32(orgId),
		Offset: int32(offset),
		Limit:  int32(limit),
	})
}

// UpdateGoodClass 更新产品类型
func (c *GoodsCase) UpdateGoodsClass(ctx context.Context, class *db.Class) (*db.Class, error) {
	key := cache.GenKey("goods", "class", "id", int32(class.ID))
	err := c.repo.ExecTx(ctx, func(queries *db.Queries) error {
		if err := c.repo.CacheUpdate(ctx, key, func(q db.Querier) error {
			if err := queries.UpdateGoodClass(ctx, &db.UpdateGoodClassParams{
				Name: sql.NullString{
					String: class.Name.String,
					Valid:  true,
				},
				Des: class.Des,
				Status: sql.NullString{
					String: class.Status.String,
					Valid:  true,
				},
				Tm: sql.NullString{
					String: class.Tm.String,
					Valid:  true,
				},
				MaterialID: class.MaterialID,
			}); err != nil {
				logger.Error(err)
				if errors.Is(err, sql.ErrNoRows) {
					return v1.ErrorGoodsClassNotFound("产品种类未找到: %v\n", class.ID)
				}
				return err
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return class, nil
}

// 获取所属企业
func (c *GoodsCase) GetOrgOfX(ctx context.Context, x string, id int32) (int32, error) {
	switch x {
	case v1.GetOrgOfXRequest_X_name[int32(v1.GetOrgOfXRequest_class)]:
		return c.repo.GetOrgOfClass(ctx, id)
	case v1.GetOrgOfXRequest_X_name[int32(v1.GetOrgOfXRequest_serial)]:
		return c.repo.GetOrgOfSerial(ctx, id)
	case v1.GetOrgOfXRequest_X_name[int32(v1.GetOrgOfXRequest_goods)]:
		return c.repo.GetOrgOfGoods(ctx, id)
	}

	return -1, fmt.Errorf("不合法的参数：%v", x)
}
