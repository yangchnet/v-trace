package contract

import (
	"math/big"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	"github.com/ethereum/go-ethereum/common"
)

type Interface interface {
	// Name 返回合约名称
	Name() (string, error)

	// Symbol 返回合约标识
	Symbol() (string, error)

	// BalanceOf 查询用户资产数量
	BalanceOf(owner evmutils.Address) (*big.Int, error)

	// OwnerOf 查询token的所有者
	OwnerOf(tokenId *big.Int) (common.Address, error)

	// SafeMint 安全的铸造一个`tokenId`，并将其转发給`to`
	// require:
	// - transId不存在
	// - caller必须为MINTER_ROLE
	SafeMint(to evmutils.Address, tokenId *big.Int, transId, digest string) error

	// TokenUri 返回 `tokenId` 令牌的统一资源标识符 (URI)。
	TokenUri(tokenId *big.Int) (string, error)

	// Burn 销毁`tokenId`
	Burn(tokenId *big.Int) error

	// DefaultAdminRole 查询默认管理员，一般为合约部署者
	DefaultAdminRole() ([32]byte, error)

	// MinterRole 查询铸币者角色成员
	MinterRole() ([32]byte, error)

	// GetRoleAdmin 返回管理角色的管理员角色
	GetRoleAdmin(role [32]byte) ([32]byte, error)

	// GrantRole 为`account`分配`role`角色
	// caller必须为管理员角色
	GrantRole(role [32]byte, account evmutils.Address) error

	// HasRole 检查`account`是否为`role`
	HasRole(role [32]byte, account evmutils.Address) (bool, error)

	// RevokeRole 撤销`account`的`role`
	// caller必须为管理员角色
	RevokeRole(role [32]byte, account evmutils.Address) error

	// TransferFrom 将tokenId从from转让到to
	// require:
	// - from不为0地址，to不为0地址
	// - tokenId必须存在且属于from
	// - 如果caller（交易发送者）不是from，
	//		那么他必须被{approve} 或 {setApprovalForAll} 授权过
	TransferFrom(from, to evmutils.Address, tokenId *big.Int, transId, digest string) error

	// SafeTransferFrom 安全的将token从`from`转让到`to`
	// require:
	// - `from`不为0地址，`to`不为0地址
	// - tokenId必须存在且属于`from`
	// - 如果caller（交易发送者）不是`from`，
	//		那么他必须被{approve} 或 {setApprovalForAll} 授权过
	SafeTransferFrom(from, to evmutils.Address, tokenId *big.Int, transId, digest string) error

	// Approve 授予to向其他用户转让tokenId的权限
	// token被转让后，授权将会被清除
	// 同一时刻只能有一个用户被授权，因此对0地址授权清除了之前的approve
	// require:
	// - caller必须为token持有者或已被授权的操作者
	// - tokenId 必须存在
	Approve(to evmutils.Address, tokenId *big.Int) error

	// SetApprovalForAll 批准或删除 `operator` 作为caller的Operator
	// Operator可以为调用者拥有的任何token调用 {transferFrom} 或 {safeTransferFrom}
	// require:
	// - `operator`不可为caller
	SetApprovalForAll(operator evmutils.Address, approved bool) error

	// 返回`operator`是否被授权管理`owner`的所有资产
	IsApprovedForAll(owner, operator evmutils.Address) (bool, error)

	// GetApproved 返回被授权处理`tokenId`的账户
	GetApproved(tokenId *big.Int) (common.Address, error)

	// Process 为token 附加一条信息
	// require:
	// - caller必需为token的owner
	// - transId不存在
	Process(tokenId *big.Int, transId, digest string) error

	// GetTransIds 获取token的transId记录
	GetTransIds(tokenId *big.Int) ([]string, error)

	// TransDigest 获取交易表单摘要
	TransDigest(transId string) (string, error)

	// BatchMint 批量铸币
	// require:
	// - tokenId 不存在
	// - transId 不存在
	// - caller必须为MINTER_ROLE
	BatchMint(to evmutils.Address, tokenIds []*big.Int, transIds []string, digest string) error

	// BatchBurn 批量销毁
	// require:
	// - caller为token的owner
	BatchBurn(tokenIds []*big.Int) error

	// BatchProcess 批量处理
	// require:
	// - caller为token的owner
	BatchProcess(tokenIds []*big.Int, transIds []string, digest string) error

	// BatchTransform 批量转帐
	// require:
	// - `from`不为0地址，`to`不为0地址
	// - tokenId必须存在且属于`from`
	// - 如果caller（交易发送者）不是`from`，
	//		那么他必须被{approve} 或 {setApprovalForAll} 授权过
	BatchTransform(from, to evmutils.Address, tokenIds []*big.Int, transIds []string, digest string) error
}
