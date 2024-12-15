package biz

import (
	"context"
	"errors"

	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/service"
	"gitee.com/qciip-icp/v-trace/pkg/fs"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewVTraceCase)

var ErrPermissionDenied = errors.New("鉴权失败, 请检查是否有权限访问资源")

var _ service.VTraceCaseInterface = (*VTraceCase)(nil)

type VTraceCase struct {
	repo VTraceRepo
	fs   fs.Interface
}

func NewVTraceCase(ctx context.Context, repo VTraceRepo, fs fs.Interface) *VTraceCase {
	return &VTraceCase{
		repo: repo,
		fs:   fs,
	}
}
