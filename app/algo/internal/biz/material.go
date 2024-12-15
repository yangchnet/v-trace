package biz

import (
	"context"
	"database/sql"

	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
)

func (c *AlgoCase) ListMaterials(ctx context.Context) ([]*db.Material, [][]*db.Model, error) {
	materials, err := c.repo.ListMaterials(ctx)
	if err != nil {
		logger.Error(err)

		return nil, nil, err
	}

	models := make([][]*db.Model, 0)
	for _, m := range materials {
		ms, err := c.repo.GetAvaiableModel(ctx, sql.NullInt32{
			Int32: m.ID,
			Valid: true,
		})
		if err != nil {
			logger.Error(err)
			continue
		}

		models = append(models, ms)
	}

	return materials, models, nil
}
