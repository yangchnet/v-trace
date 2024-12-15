package service

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
)

type AlgoCaseInterface interface {
	ListModels(ctx context.Context) ([]*db.Model, error)
	ListMaterials(ctx context.Context) ([]*db.Material, [][]*db.Model, error)
	Predict(ctx context.Context, modelName string, data []byte) (*db.Material, error)
}
