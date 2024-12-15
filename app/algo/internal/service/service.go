package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAlgoService)

type AlgoService struct {
	v1.UnimplementedAlgoServiceServer

	cas AlgoCaseInterface
}

func (s *AlgoService) mustEmbedUnimplementedAlgoServiceServer() {
	panic("not implemented") // TODO: Implement
}

func NewAlgoService(ctx context.Context, cas AlgoCaseInterface) *AlgoService {
	return &AlgoService{
		cas: cas,
	}
}

func (s *AlgoService) GetDomain() string {
	return "AlgoService"
}
