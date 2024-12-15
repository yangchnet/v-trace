package biz

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/pkg/algo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAlgoCase)

type AlgoCase struct {
	repo   AlgoRepo
	client *algo.AlgoClient
}

func NewAlgoCase(ctx context.Context, repo AlgoRepo, algoClient *algo.AlgoClient) *AlgoCase {
	return &AlgoCase{
		repo:   repo,
		client: algoClient,
	}
}
