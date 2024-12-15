package data

import (
	"context"
	"database/sql"
	"encoding/json"

	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/gotools"
)

func (d *Data) GetMaterialByID(ctx context.Context, id int32) (*db.Material, error) {
	ck := cache.GenKey("algo", "material", "id", id)

	mBytes, ok := d.cache.Get(ctx, ck)
	if ok {
		var m db.Material
		if err := json.Unmarshal(mBytes.([]byte), &m); err == nil {
			return &m, nil
		}
	}

	m, err := d.Store.GetMaterialByID(ctx, id)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	gotools.Go(func() {
		_ = cache.MarshaledAndCached(ctx, d.cache, ck, m)
	})

	return m, nil
}

// GetModelByName 根据模型名获取模型数据.
func (d *Data) GetModelByName(ctx context.Context, name sql.NullString) (*db.Model, error) {
	ck := cache.GenKey("algo", "model", "name", name.String)
	mBytes, ok := d.cache.Get(ctx, ck)
	if ok {
		var m db.Model
		if err := json.Unmarshal(mBytes.([]byte), &m); err == nil {
			return &m, nil
		}
	}

	m, err := d.Store.GetModelByName(ctx, name)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	gotools.Go(func() {
		_ = cache.MarshaledAndCached(ctx, d.cache, ck, m)
	})

	return m, err
}
