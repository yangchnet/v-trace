package biz

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/circ/internal/data"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCircCase)

type CircCase struct {
	repo CircRepo
}

type CircRepo interface {
	data.Store
}

func NewCircCase(ctx context.Context, repo CircRepo) *CircCase {
	return &CircCase{
		repo: repo,
	}
}
