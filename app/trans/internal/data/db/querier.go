// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	// CreateTrans 创建交易记录
	CreateTrans(ctx context.Context, arg *CreateTransParams) (int64, error)
	// 根据transId获取记录
	GetTransByTransId(ctx context.Context, transid string) (*TransRecord, error)
	// 根据transId更新交易状态
	UpdateTrans(ctx context.Context, arg *UpdateTransParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
