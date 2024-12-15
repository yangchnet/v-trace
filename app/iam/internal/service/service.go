package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewIamService)

type IamService struct {
	v1.UnimplementedIamServiceServer

	cas IamCaseInterface
}

var _ v1.IamServiceServer = (*IamService)(nil)

func (s *IamService) mustEmbedUnimplementedIamServiceServer() {
	panic("not implemented") // TODO: Implement
}

func NewIamService(ctx context.Context, cas IamCaseInterface) *IamService {
	return &IamService{cas: cas}
}

func (s *IamService) GetDomain() string {
	return "IamService"
}
