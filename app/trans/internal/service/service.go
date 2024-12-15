package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/data/db"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewTransService)

type TransService struct {
	v1.UnimplementedTransServiceServer

	cas TransCaseInterface
}

func NewTransService(ctx context.Context, cas TransCaseInterface) *TransService {
	return &TransService{
		cas: cas,
	}
}

type TransCaseInterface interface {
	// Call调用合约
	Name(ctx context.Context, caller *v1.Identity, transId string) error
	// mint
	Mint(ctx context.Context, caller *v1.Identity, to *v1.Identity, tokenId int64, transId, digest string) error
	// Burn
	Burn(ctx context.Context, caller *v1.Identity, tokenId int64, transId string) error
	// Transfer
	Transfer(ctx context.Context, caller *v1.Identity, from *v1.Identity, to *v1.Identity, tokenId int64, transId, digest string) error
	// SetApproveForAll
	SetApproveForAll(ctx context.Context, caller, operator *v1.Identity, approved bool, transId string) error
	// GrantRole
	GrantRole(ctx context.Context, caller, account *v1.Identity, transId string) error
	// GetTransRecords
	GetTransRecord(ctx context.Context, transId string) (*db.TransRecord, error)
	// UpdateTrans
	UpdateTrans(ctx context.Context, transId, txHash string, success bool) error
	// AddPrecess
	Process(ctx context.Context, caller *v1.Identity, tokenId int64, transId, digest string) error
	BatchMint(ctx context.Context, caller *v1.Identity, to *v1.Identity, tokenIds []int64, transIds []string, digest string) error
	BatchBurn(ctx context.Context, caller *v1.Identity, tokenIds []int64, transIds []string) error
	BatchProcess(ctx context.Context, caller *v1.Identity, tokenIds []int64, transIds []string, digest string) error
	BatchTransform(ctx context.Context, caller *v1.Identity, from, to *v1.Identity, tokenIds []int64, transIds []string, digest string) error
}

func (s *TransService) GetDomain() string {
	return "trans-service"
}
