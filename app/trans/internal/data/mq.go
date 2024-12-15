package data

import (
	"context"

	"gitee.com/qciip-icp/v-trace/pkg/constants"
)

func (r *Data) MqAdd(ctx context.Context, payload []byte) error {
	return r.mq.Add(ctx, constants.MqStream, constants.MqEmptyId, constants.MqKey, payload)
}
