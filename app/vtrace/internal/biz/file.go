package biz

import (
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// 存储文件
func (uc *VTraceCase) Store(ctx context.Context, content []byte, metadata map[string]string) (string, error) {
	key := uuid.New().String()

	if _, err := uc.fs.Store(content, key, metadata, nil); err != nil {
		logger.Errorf("failed store file: %v\n", err)

		return "", err
	}

	return uc.fs.Path(key)
}
