package biz

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	cmSdk "chainmaker.org/chainmaker/sdk-go/v2"
	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/app/pkg/task"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/ethereum/go-ethereum/crypto"
)

func (uc *TransCase) Name(ctx context.Context, caller *v1.Identity, transId string) error {
	methodName := "name"

	// 1. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller)
	if err != nil {
		return err
	}

	// 2. 写入队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 3. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{}); err != nil {
		return err
	}

	return nil
}

// mint.
func (uc *TransCase) Mint(ctx context.Context, caller *v1.Identity, to *v1.Identity, tokenId int64, transId, digest string) error {
	methodName := contract.MethodSafeMint

	toAddr, err := addrFromCert(to.GetCert())
	if err != nil {
		logger.Error(err)

		return err
	}

	// 1. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller, toAddr, big.NewInt(tokenId), transId, digest)
	if err != nil {
		return err
	}

	// 2. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 3. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"to":      to.GetUsername().GetValue(),
		"tokenId": tokenId,
		"transId": transId,
		"digest":  digest,
	}); err != nil {
		return err
	}

	return nil
}

// Burn.
func (uc *TransCase) Burn(ctx context.Context, caller *v1.Identity, tokenId int64, transId string) error {
	methodName := "burn"

	// 1. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller, big.NewInt(tokenId))
	if err != nil {
		return err
	}

	// 2. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 3. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"tokenId": tokenId,
	}); err != nil {
		return err
	}

	return nil
}

// Transfer.
func (uc *TransCase) Transfer(ctx context.Context, caller *v1.Identity, from *v1.Identity, to *v1.Identity, tokenId int64, transId, digest string) error {
	methodName := contract.MethodSafeTransferFrom

	// 1. 解析证书获取参数
	fromAddr, err := addrFromCert(from.GetCert())
	if err != nil {
		logger.Error(err)

		return err
	}
	toAddr, err := addrFromCert(to.GetCert())
	if err != nil {
		logger.Error(err)

		return err
	}

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller, fromAddr, toAddr, big.NewInt(tokenId), transId, digest)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"from":    from.GetUsername().GetValue(),
		"to":      to.GetUsername().GetValue(),
		"tokenId": tokenId,
		"transId": transId,
		"digest":  digest,
	}); err != nil {
		return err
	}

	return nil
}

// SetApproveForAll.
func (uc *TransCase) SetApproveForAll(ctx context.Context, caller *v1.Identity, entrusted *v1.Identity, approved bool, transId string) error {
	methodName := "setApprovalForAll"

	// 1. 解析证书获取参数
	entrustedAddr, err := addrFromCert(entrusted.GetCert())
	if err != nil {
		logger.Error(err)

		return err
	}

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller, entrustedAddr, approved)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"entrusted": entrusted.GetUsername().GetValue(),
		"approved":  approved,
	}); err != nil {
		return err
	}

	return nil
}

// GrantRole.
func (uc *TransCase) GrantRole(ctx context.Context, caller *v1.Identity, account *v1.Identity, transId string) error {
	methodName := "grantRole"

	// 1. 获取参数
	accountAddr, err := addrFromCert(account.GetCert())
	if err != nil {
		logger.Error(err)

		return err
	}

	minterRoleByte := crypto.Keccak256([]byte("MINTER_ROLE"))
	var minterRole [32]byte
	copy(minterRole[:], minterRoleByte)

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller, minterRole, accountAddr)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"role":    minterRole,
		"account": account.GetUsername().GetValue(),
	}); err != nil {
		return err
	}

	return nil
}

// Precess.
func (uc *TransCase) Process(ctx context.Context, caller *v1.Identity, tokenId int64, transId, digest string) error {
	methodName := contract.MethodAddProcess

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transId, caller, big.NewInt(tokenId), transId, digest)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transId, uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"tokenId": tokenId,
		"transId": transId,
		"digest":  digest,
	}); err != nil {
		return err
	}

	return nil
}

// from *v1.Identity, to *v1.Identity, tokenId int64, transId, digest string
// BatchMint 批量铸币
// require:
// - tokenId 不存在
// - transId 不存在
// - caller必须为MINTER_ROLE
func (uc *TransCase) BatchMint(ctx context.Context, caller *v1.Identity, to *v1.Identity, tokenIds []int64, transIds []string, digest string) error {
	methodName := contract.MethodBatchMint

	bigTokenIds := make([]*big.Int, 0)
	for _, token := range tokenIds {
		bigTokenIds = append(bigTokenIds, big.NewInt(token))
	}

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transIds[0], caller, bigTokenIds, transIds, digest)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transIds[0], uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"to":       to.GetUsername().GetValue(),
		"tokenIds": tokenIds,
		"transIds": transIds,
		"digest":   digest,
	}); err != nil {
		return err
	}

	return nil
}

// BatchBurn 批量销毁
// require:
// - caller为token的owner
func (uc *TransCase) BatchBurn(ctx context.Context, caller *v1.Identity, tokenIds []int64, transIds []string) error {
	methodName := contract.MethodBatchMint

	bigTokenIds := make([]*big.Int, 0)
	for _, token := range tokenIds {
		bigTokenIds = append(bigTokenIds, big.NewInt(token))
	}

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transIds[0], caller, bigTokenIds)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transIds[0], uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"tokenIds": tokenIds,
		"transIds": transIds,
	}); err != nil {
		return err
	}

	return nil
}

// BatchProcess 批量处理
// require:
// - caller为token的owner
func (uc *TransCase) BatchProcess(ctx context.Context, caller *v1.Identity, tokenIds []int64, transIds []string, digest string) error {
	methodName := contract.MethodBatchProcess

	bigTokenIds := make([]*big.Int, 0)
	for _, token := range tokenIds {
		bigTokenIds = append(bigTokenIds, big.NewInt(token))
	}

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transIds[0], caller, bigTokenIds, transIds, digest)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transIds[0], uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"tokenIds": tokenIds,
		"transIds": transIds,
		"digest":   digest,
	}); err != nil {
		return err
	}

	return nil
}

// BatchTransform 批量转帐
// require:
//   - `from`不为0地址，`to`不为0地址
//   - tokenId必须存在且属于`from`
//   - 如果caller（交易发送者）不是`from`，
//     那么他必须被{approve} 或 {setApprovalForAll} 授权过
func (uc *TransCase) BatchTransform(ctx context.Context, caller *v1.Identity, from, to *v1.Identity, tokenIds []int64, transIds []string, digest string) error {
	methodName := contract.MethodBatchProcess

	bigTokenIds := make([]*big.Int, 0)
	for _, token := range tokenIds {
		bigTokenIds = append(bigTokenIds, big.NewInt(token))
	}

	// 2. 打包交易
	payloadBytes, err := uc.pack(ctx, methodName, transIds[0], caller, from, to, bigTokenIds, transIds, digest)
	if err != nil {
		return err
	}

	// 3. 发送到队列
	if err := uc.repo.MqAdd(ctx, payloadBytes); err != nil {
		logger.Error(err)

		return err
	}

	// 4. 记录到数据库
	if err := uc.repo.CreateTransWithParams(ctx, transIds[0], uc.ContractData.RawContractName, caller.GetUsername().GetValue(), methodName, map[string]any{
		"form":     from.GetUsername().GetValue(),
		"to":       to.GetUsername().GetValue(),
		"tokenIds": tokenIds,
		"transIds": transIds,
		"digest":   digest,
	}); err != nil {
		return err
	}

	return nil
}

func (uc *TransCase) pack(ctx context.Context, methodName, transId string, call *v1.Identity, args ...interface{}) ([]byte, error) {
	packData, err := uc.ContractData.Abi.Pack(methodName, args...)
	if err != nil {
		logger.Errorf("error pack transaciont data: %+v", err)

		return nil, err
	}

	dataString := hex.EncodeToString(packData)
	method := dataString[0:8]

	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}

	payload := task.TransTask{
		TaskId:       transId,
		Crt:          call.Cert,
		Pk:           call.Key,
		TlsCrt:       call.TlsCert,
		TlsPk:        call.TlsKey,
		ContractName: uc.ContractData.ContractName,
		Method:       method,
		MethodName:   methodName,
		Kv:           kvs,
		OrgId:        constants.OrgId,
		ChainId:      constants.ChainId,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Errorf("error marshal payload: %+v", err)

		return nil, err
	}

	return payloadBytes, nil
}

func addrFromCert(cert []byte) (evmutils.Address, error) {
	addrStr, err := cmSdk.GetEVMAddressFromCertBytes(cert)
	if err != nil {
		return evmutils.Address{}, err
	}
	addr, err := evmutils.HexToAddress(addrStr)
	if err != nil {
		return evmutils.Address{}, err
	}

	return addr, nil
}
