package data

import (
	"context"

	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
)

type CallContractParams struct {
	MethodName string
	Caller     *v1.Identity
	From       *v1.Identity
	To         *v1.Identity
	TokenId    []int64
	TransId    []string
	Digest     string
}

func (uc *Data) CallContract(ctx context.Context, params *CallContractParams) error {
	var err error
	switch params.MethodName {
	case contract.MethodSafeMint:
		_, err = uc.Trans().Mint(ctx, &v1.MintRequest{
			Caller:  params.Caller,
			To:      params.From,
			TokenId: params.TokenId[0],
			TransId: params.TransId[0],
			Digest:  params.Digest,
		})
	case contract.MethodBurn:
		_, err = uc.Trans().Burn(
			ctx,
			&v1.BurnRequest{
				Caller:  params.Caller,
				TokenId: params.TokenId[0],
				TransId: params.TransId[0],
			},
		)
	case contract.MethodTransferFrom:
		_, err = uc.Trans().Transfer(
			ctx,
			&v1.TransferRequest{
				Caller:  params.Caller,
				From:    params.From,
				To:      params.To,
				TokenId: params.TokenId[0],
				TransId: params.TransId[0],
				Digest:  params.Digest,
			},
		)
	case contract.MethodApprove:
		_, err = uc.Trans().Approve(
			ctx,
			&v1.ApproveRequest{
				Caller:  params.Caller,
				To:      params.To,
				TokenId: params.TokenId[0],
			},
		)
	case contract.MethodAddProcess:
		_, err = uc.Trans().AddProcess(
			ctx,
			&v1.AddProcessRequest{
				Caller:  params.Caller,
				TokenId: params.TokenId[0],
				TransId: params.TransId[0],
				Digest:  params.Digest,
			},
		)
	case contract.MethodGrantRole:
		_, err = uc.Trans().GrantRole(
			ctx,
			&v1.GrantRoleRequest{
				Granter: params.Caller,
				Account: params.To,
				TransId: params.TransId[0],
			},
		)
	case contract.MethodBatchMint:
		_, err = uc.Trans().BatchMint(ctx,
			&v1.BatchMintRequest{
				Caller:   params.Caller,
				To:       params.To,
				TokenIds: params.TokenId,
				TransIds: params.TransId,
				Digest:   params.Digest,
			},
		)
	case contract.MethodBatchBurn:
		_, err = uc.Trans().BatchBurn(ctx,
			&v1.BatchBurnRequest{
				Caller:   params.Caller,
				TokenIds: params.TokenId,
			},
		)
	case contract.MethodBatchProcess:
		_, err = uc.Trans().BatchProcess(ctx,
			&v1.BatchProcessRequest{
				Caller:   params.Caller,
				TokenIds: params.TokenId,
				TransIds: params.TransId,
				Digest:   params.Digest,
			},
		)
	case contract.MethodBatchTransform:
		_, err = uc.Trans().BatchTransform(ctx,
			&v1.BatchTransformRequest{
				Caller:   params.Caller,
				From:     params.From,
				To:       params.To,
				TokenIds: params.TokenId,
				TransIds: params.TransId,
				Digest:   params.Digest,
			},
		)
	}

	logger.Error(err)

	return err
}

// 根据transId获取交易信息.
func (r *Data) GetTxByTransId(ctx context.Context, transId string) (*transV1.TransRecord, error) {
	tx, err := r.Trans().GetTrans(ctx, &v1.GetTransRequest{
		TransId: transId,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return tx.Trans, nil
}
