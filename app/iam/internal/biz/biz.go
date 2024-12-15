package biz

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/iam/internal/data"
	"gitee.com/qciip-icp/v-trace/pkg/token"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewIamCase)

type IamRepo interface {
	data.Interface
}

type IamCase struct {
	repo       IamRepo
	tokenMaker token.Maker
}

func NewIamCase(
	ctx context.Context,
	iamRepo IamRepo,
	tokenMaker token.Maker,
) *IamCase {
	return &IamCase{
		repo:       iamRepo,
		tokenMaker: tokenMaker,
	}
}
