package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

func (s *TransService) Name(ctx context.Context, req *v1.NameRequest) (*v1.NameResponse, error) {
	caller := req.GetCaller()
	transId := req.GetTransId()

	if err := s.cas.Name(ctx, caller, transId); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.NameResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) Mint(ctx context.Context, req *v1.MintRequest) (*v1.MintResponse, error) {
	callerIdentity := req.GetCaller()
	tokenId := req.GetTokenId()
	transId := req.GetTransId()
	toIdentity := req.GetTo()

	if err := s.cas.Mint(ctx, callerIdentity, toIdentity, tokenId, transId, req.GetDigest()); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.MintResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) Burn(ctx context.Context, req *v1.BurnRequest) (*v1.BurnResponse, error) {
	identity := req.GetCaller()
	tokenId := req.GetTokenId()
	transId := req.GetTransId()

	if err := s.cas.Burn(ctx, identity, tokenId, transId); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.BurnResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) Transfer(ctx context.Context, req *v1.TransferRequest) (*v1.TransferResponse, error) {
	operator := req.GetCaller()
	from := req.GetFrom()
	to := req.GetTo()
	tokenId := req.GetTokenId()
	transId := req.GetTransId()

	if err := s.cas.Transfer(ctx, operator, from, to, tokenId, transId, req.GetDigest()); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.TransferResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) SetApproveForAll(ctx context.Context, req *v1.SetApproveForAllRequest) (*v1.SetApproveForAllResponse, error) {
	caller := req.GetCaller()
	operator := req.GetOperator()
	approved := req.GetApproved()
	transId := req.GetTransId()

	if err := s.cas.SetApproveForAll(ctx, caller, operator, approved, transId); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.SetApproveForAllResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) GrantRole(ctx context.Context, req *v1.GrantRoleRequest) (*v1.GrantRoleResponse, error) {
	granter := req.GetGranter()
	account := req.GetAccount()
	transId := req.GetTransId()

	if err := s.cas.GrantRole(ctx, granter, account, transId); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.GrantRoleResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) Symbol(_ context.Context, _ *v1.SymbolRequest) (*v1.SymbolResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *TransService) GetTrans(ctx context.Context, req *v1.GetTransRequest) (*v1.GetTransResponse, error) {
	trans, err := s.cas.GetTransRecord(ctx, req.GetTransId())
	if err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.GetTransResponse{
		Trans: Record2Proto(trans),
	}, nil
}

func (s *TransService) UpdateTrans(ctx context.Context, req *v1.UpdateTransRequest) (*v1.UpdateTransResponse, error) {
	if err := s.cas.UpdateTrans(ctx, req.GetTransId(), req.GetTxHash(), req.GetSuccess()); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.UpdateTransResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) AddProcess(ctx context.Context, req *v1.AddProcessRequest) (*v1.AddProcessResponse, error) {
	if err := s.cas.Process(
		ctx,
		req.GetCaller(),
		req.GetTokenId(),
		req.GetTransId(),
		req.GetDigest(),
	); err != nil {
		return nil, verr.Error(s, err)
	}

	return &v1.AddProcessResponse{
		Result: v1.OperationResult_SUCCESS,
	}, nil
}

func (s *TransService) mustEmbedUnimplementedTransServiceServer() {
	panic("not implemented") // TODO: Implement
}
