package task

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
)

type TransTask struct {
	// 任务id -> transId
	TaskId string

	// 用户证书
	Crt []byte `json:"crt"`

	// 用户密钥
	Pk []byte `json:"pk"`

	// 用户tls证书
	TlsCrt []byte `json:"tls_crt"`

	// 用户tls密钥
	TlsPk []byte `json:"tls_pk"`

	// 合约名，16进制
	ContractName string `json:"contractName"`

	// 函数名，字符串
	MethodName string `json:"method"`

	// 16进制标识的函数名
	Method string `json:"methodName"`

	//
	Kv []*common.KeyValuePair `json:"kv"`

	// 组织id
	OrgId string `json:"orgId"`

	// 链id
	ChainId string `json:"chainId"`
}
