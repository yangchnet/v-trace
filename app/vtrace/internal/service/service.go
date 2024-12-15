package service

import (
	"context"

	"gitee.com/qciip-icp/v-trace/pkg/token"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewVTraceService)

type VTraceService struct {
	v1.UnimplementedVTraceInterfaceServer

	cas VTraceCaseInterface

	tokenMaker token.Maker
}

var _ v1.VTraceInterfaceServer = (*VTraceService)(nil)

func NewVTraceService(ctx context.Context, cas VTraceCaseInterface) *VTraceService {
	return &VTraceService{
		cas: cas,
	}
}

func (s *VTraceService) GetDomain() string {
	return "vtrace-app"
}
