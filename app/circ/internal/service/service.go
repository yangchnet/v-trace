package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCircService)

type CircService struct {
	v1.UnimplementedCircServiceServer

	cas CircCaseInterface
}

func (s *CircService) mustEmbedUnimplementedCircServiceServer() {
	panic("not implemented") // TODO: Implement
}

func NewCircService(ctx context.Context, cas CircCaseInterface) *CircService {
	return &CircService{
		cas: cas,
	}
}

func (s *CircService) GetDomain() string {
	return "CircService"
}
