package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/google/wire"
)

var ContractProvider = wire.NewSet(NewVTraceContract, NewContractData)

var CallAction map[string]string = map[string]string{
	circV1.CircType_name[int32(circV1.CircType_produce)]:  MethodSafeMint,
	circV1.CircType_name[int32(circV1.CircType_process)]:  MethodAddProcess,
	circV1.CircType_name[int32(circV1.CircType_transfer)]: MethodSafeTransferFrom,
	circV1.CircType_name[int32(circV1.CircType_exam)]:     MethodAddProcess,
}

var BatchCallAction map[string]string = map[string]string{
	circV1.CircType_name[int32(circV1.CircType_produce)]:  MethodBatchMint,
	circV1.CircType_name[int32(circV1.CircType_process)]:  MethodBatchProcess,
	circV1.CircType_name[int32(circV1.CircType_transfer)]: MethodBatchTransform,
	circV1.CircType_name[int32(circV1.CircType_exam)]:     MethodBatchProcess,
}

var (
	MethodName              = "name"
	MethodSymbol            = "symbol"
	MethodBalanceOf         = "balanceOf"
	MethodOwnerOf           = "ownerOf"
	MethodSafeMint          = "safeMint"
	MethodTokenUri          = "tokenURI"
	MethodBurn              = "burn"
	MethodDefaultAdminRole  = "DEFAULT_ADMIN_ROLE"
	MethodMinterRole        = "MINTER_ROLE"
	MethodGetRoleAdmin      = "getRoleAdmin"
	MethodGrantRole         = "grantRole"
	MethodHasRole           = "hasRole"
	MethodRevokeRole        = "revokeRole"
	MethodTransferFrom      = "transferFrom"
	MethodSafeTransferFrom  = "safeTransferFrom"
	MethodApprove           = "approve"
	MethodSetApprovalForAll = "setApprovalForAll"
	MethodIsApprovedForAll  = "isApprovedForAll"
	MethodGetApproved       = "getApproved"
	MethodAddProcess        = "process"
	MethodGetTransIds       = "getTransIds"
	MethodTransDigest       = "transDigest"
	MethodBatchMint         = "batchMint"
	MethodBatchBurn         = "batchBurn"
	MethodBatchProcess      = "batchProcess"
	MethodBatchTransform    = "batchTransform"
)

type VTraceContract struct {
	Client         *sdk.ChainClient
	Abi            *abi.ABI
	BinCode        []byte
	WithSyncResult bool
	ContractName   string
}

type ContractData struct {
	Abi             *abi.ABI
	BinCode         []byte
	ContractName    string
	RawContractName string
}

type ContractCallResp struct {
	TxHash string
	Res    interface{}
}

func NewRawVTraceContract(contractData *ContractData, withSyncResult bool, name string) *VTraceContract {
	return &VTraceContract{
		Abi:            contractData.Abi,
		BinCode:        contractData.BinCode,
		WithSyncResult: withSyncResult,
		ContractName:   name,
	}
}

func (c *VTraceContract) WithClient(ctx context.Context, client *sdk.ChainClient) {
	c.Client = client
}

func NewVTraceContract(client *sdk.ChainClient, contractData *ContractData, withSyncResult bool, name string) *VTraceContract {
	return &VTraceContract{
		Client:         client,
		Abi:            contractData.Abi,
		BinCode:        contractData.BinCode,
		WithSyncResult: withSyncResult,
		ContractName:   name,
	}
}

func NewContractData(abiPath, binPath, contractName string) (*ContractData, error) {
	var data ContractData = ContractData{}

	abiJson, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	myAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		return nil, err
	}

	data.Abi = &myAbi

	binCode, err := ioutil.ReadFile(binPath)
	if err != nil {
		return nil, err
	}

	data.BinCode = binCode

	data.ContractName = CalcContractName(contractName)

	data.RawContractName = contractName

	return &data, nil
}

var _ Interface = (*VTraceContract)(nil)

// Name 返回合约名称.
func (c *VTraceContract) Name() (string, error) {
	resp, err := c._callMethod(MethodName)
	if err != nil {
		return "", err
	}

	name, err := c.Abi.Unpack(MethodName, resp.ContractResult.Result)
	if err != nil {
		return "", err
	}

	return name[0].(string), nil
}

// Symbol 返回合约标识.
func (c *VTraceContract) Symbol() (string, error) {
	resp, err := c._callMethod(MethodSymbol)
	if err != nil {
		return "", err
	}

	symbol, err := c.Abi.Unpack(MethodSymbol, resp.ContractResult.Result)
	if err != nil {
		return "", err
	}

	return symbol[0].(string), nil
}

// BalanceOf 查询用户资产数量.
func (c *VTraceContract) BalanceOf(owner evmutils.Address) (*big.Int, error) {
	resp, err := c._callMethod(MethodBalanceOf, owner)
	if err != nil {
		return nil, err
	}

	balance, err := c.Abi.Unpack(MethodBalanceOf, resp.ContractResult.Result)
	if err != nil {
		return nil, err
	}

	return balance[0].(*big.Int), nil
}

// OwnerOf 查询token的所有者.
func (c *VTraceContract) OwnerOf(tokenId *big.Int) (ethCommon.Address, error) {
	methodName := MethodOwnerOf

	resp, err := c._callMethod(methodName, tokenId)
	if err != nil {
		return ethCommon.Address{}, err
	}

	owner, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return ethCommon.Address{}, err
	}

	return owner[0].(ethCommon.Address), nil
}

// SafeMint 安全的铸造一个`tokenId`，并将其转发給`to`.
func (c *VTraceContract) SafeMint(to evmutils.Address, tokenId *big.Int, transId, digest string) error {
	methodName := MethodSafeMint

	resp, err := c._callMethod(methodName, to, tokenId, transId, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// TokenUri 返回 `tokenId` 令牌的统一资源标识符 (URI)。
func (c *VTraceContract) TokenUri(tokenId *big.Int) (string, error) {
	methodName := MethodTokenUri

	resp, err := c._callMethod(methodName, tokenId)
	if err != nil {
		return "", err
	}

	uri, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return "", err
	}

	return uri[0].(string), nil
}

// Burn 销毁`tokenId`.
func (c *VTraceContract) Burn(tokenId *big.Int) error {
	methodName := MethodBurn

	resp, err := c._callMethod(methodName, tokenId)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// DefaultAdminRole 查询默认管理员，一般为合约部署者.
func (c *VTraceContract) DefaultAdminRole() ([32]byte, error) {
	methodName := MethodDefaultAdminRole

	resp, err := c._callMethod(methodName)
	if err != nil {
		return [32]byte{}, err
	}

	admin, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return [32]byte{}, err
	}

	return admin[0].([32]byte), nil
}

// MinterRole 查询铸币者角色.
func (c *VTraceContract) MinterRole() ([32]byte, error) {
	methodName := MethodMinterRole

	resp, err := c._callMethod(methodName)
	if err != nil {
		return [32]byte{}, err
	}

	minters, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return [32]byte{}, err
	}

	return minters[0].([32]byte), nil
}

// GetRoleAdmin 返回管理角色的管理员角色.
func (c *VTraceContract) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	methodName := MethodGetRoleAdmin

	resp, err := c._callMethod(methodName, role)
	if err != nil {
		return [32]byte{}, err
	}

	admin, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return [32]byte{}, err
	}

	return admin[0].([32]byte), nil
}

// GrantRole 为`account`分配`role`角色
// caller必须为管理员角色.
func (c *VTraceContract) GrantRole(role [32]byte, account evmutils.Address) error {
	methodName := MethodGrantRole

	resp, err := c._callMethod(methodName, role, account)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// HasRole 检查`account`是否为`role`.
func (c *VTraceContract) HasRole(role [32]byte, account evmutils.Address) (bool, error) {
	methodName := MethodHasRole

	resp, err := c._callMethod(methodName, role, account)
	if err != nil {
		return false, err
	}

	is, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return false, err
	}

	return is[0].(bool), nil
}

// RevokeRole 撤销`account`的`role`
// caller必须为管理员角色.
func (c *VTraceContract) RevokeRole(role [32]byte, account evmutils.Address) error {
	methodName := MethodRevokeRole

	resp, err := c._callMethod(methodName, role, account)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// TransferFrom 将tokenId从from转让到to
// require:
//   - from不为0地址，to不为0地址
//   - tokenId必须存在且属于from
//   - 如果caller（交易发送者）不是from，
//     那么他必须被{approve} 或 {setApprovalForAll} 授权过
func (c *VTraceContract) TransferFrom(from evmutils.Address, to evmutils.Address, tokenId *big.Int, transId, digest string) error {
	methodName := MethodTransferFrom

	resp, err := c._callMethod(methodName, from, to, tokenId, transId, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// SafeTransferFrom 安全的将token从`from`转让到`to`
// require:
//   - `from`不为0地址，`to`不为0地址
//   - tokenId必须存在且属于`from`
//   - 如果caller（交易发送者）不是`from`，
//     那么他必须被{approve} 或 {setApprovalForAll} 授权过
func (c *VTraceContract) SafeTransferFrom(from evmutils.Address, to evmutils.Address, tokenId *big.Int, transId, digest string) error {
	methodName := MethodSafeTransferFrom

	resp, err := c._callMethod(methodName, from, to, tokenId, transId, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// Approve 授予to向其他用户转让tokenId的权限
// token被转让后，授权将会被清除
// 同一时刻只能有一个用户被授权，因此对0地址授权清除了之前的approve
// require:
// - caller必须为token持有者或已被授权的操作者
// - tokenId 必须存在.
func (c *VTraceContract) Approve(to evmutils.Address, tokenId *big.Int) error {
	methodName := MethodApprove

	resp, err := c._callMethod(methodName, to, tokenId)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// SetApprovalForAll 批准或删除 `operator` 作为caller的Operator
// Operator可以为调用者拥有的任何token调用 {transferFrom} 或 {safeTransferFrom}
// require:
// - `operator`不可为caller.
func (c *VTraceContract) SetApprovalForAll(operator evmutils.Address, approved bool) error {
	methodName := MethodSetApprovalForAll

	resp, err := c._callMethod(methodName, operator, approved)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// 返回`operator`是否被授权管理`owner`的所有资产.
func (c *VTraceContract) IsApprovedForAll(owner evmutils.Address, operator evmutils.Address) (bool, error) {
	methodName := MethodIsApprovedForAll

	resp, err := c._callMethod(methodName, owner, operator)
	if err != nil {
		return false, err
	}

	is, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return false, err
	}

	return is[0].(bool), nil
}

// GetApproved 返回被授权处理`tokenId`的账户.
func (c *VTraceContract) GetApproved(tokenId *big.Int) (ethCommon.Address, error) {
	methodName := MethodGetApproved

	resp, err := c._callMethod(methodName, tokenId)
	if err != nil {
		return ethCommon.Address{}, err
	}

	addr, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return ethCommon.Address{}, err
	}

	return addr[0].(ethCommon.Address), nil
}

// Process 为token 附加一条信息.
func (c *VTraceContract) Process(tokenId *big.Int, transId, digest string) error {
	methodName := MethodAddProcess

	resp, err := c._callMethod(methodName, tokenId, transId, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// GetTransIds 获取token的transId记录
func (c *VTraceContract) GetTransIds(tokenId *big.Int) ([]string, error) {
	methodName := MethodGetTransIds

	resp, err := c._callMethod(methodName, tokenId)
	if err != nil {
		return nil, err
	}

	transIds, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return nil, err
	}

	return transIds[0].([]string), nil
}

// TransDigest 获取交易表单摘要
func (c *VTraceContract) TransDigest(transId string) (string, error) {
	methodName := MethodGetTransIds

	resp, err := c._callMethod(methodName, transId)
	if err != nil {
		return "", err
	}

	digest, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return "", err
	}

	return digest[0].(string), nil
}

// BatchMint 批量铸币
// require:
// - tokenId 不存在
// - transId 不存在
// - caller必须为MINTER_ROLE
func (c *VTraceContract) BatchMint(to evmutils.Address, tokenIds []*big.Int, transIds []string, digest string) error {
	methodName := MethodBatchMint

	resp, err := c._callMethod(methodName, to, tokenIds, transIds, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// BatchBurn 批量销毁
// require:
// - caller为token的owner
func (c *VTraceContract) BatchBurn(tokenIds []*big.Int) error {
	methodName := MethodBatchBurn

	resp, err := c._callMethod(methodName, tokenIds)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// BatchProcess 批量处理
// require:
// - caller为token的owner
func (c *VTraceContract) BatchProcess(tokenIds []*big.Int, transIds []string, digest string) error {
	methodName := MethodBatchProcess

	resp, err := c._callMethod(methodName, tokenIds, transIds, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
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
func (c *VTraceContract) BatchTransform(from, to evmutils.Address, tokenIds []*big.Int, transIds []string, digest string) error {
	methodName := MethodBatchTransform

	resp, err := c._callMethod(methodName, from, to, tokenIds, transIds, digest)
	if err != nil {
		return err
	}

	_, err = c.Abi.Unpack(methodName, resp.ContractResult.Result)
	if err != nil {
		return err
	}

	return nil
}

// func (c *VTraceContract) Whoami() (ethCommon.Address, error) {
// 	methodName := "whoami"

// 	resp, err := c._callMethod(methodName)
// 	if err != nil {
// 		return ethCommon.Address{}, err
// 	}

// 	owner, err := c.Abi.Unpack(methodName, resp.ContractResult.Result)
// 	if err != nil {
// 		return ethCommon.Address{}, err
// 	}

// 	return owner[0].(ethCommon.Address), nil
// }

func (c *VTraceContract) _callMethod(methodName string, args ...interface{}) (*common.TxResponse, error) {
	packData, err := c.Abi.Pack(methodName, args...)
	if err != nil {
		logger.Error(err)

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

	resp, err := c.Client.InvokeContract(
		CalcContractName(c.ContractName),
		method,
		"",
		kvs,
		-1,
		c.WithSyncResult,
	)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return nil, fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.ContractResult.Message)
	}

	return resp, nil
}
