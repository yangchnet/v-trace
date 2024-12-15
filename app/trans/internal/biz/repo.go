package biz

import (
	"context"

	"gitee.com/qciip-icp/v-trace/app/trans/internal/data"
)

type TransRepo interface {
	MqAdd(ctx context.Context, payload []byte) error
	CreateTransWithParams(ctx context.Context, transId, contractName, sender, method string, argsMap map[string]any) error

	data.Store
}
