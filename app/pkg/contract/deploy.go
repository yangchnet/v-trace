//go:build exclude

package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// call: go run app/pkg/contract/deploy.go app/pkg/contract/config/
func main() {
	if len(os.Args) <= 1 {
		panic("need config file path")
	}
	configFilePath := os.Args[1]
	DeployContractWithConfFilePath(configFilePath)
}

func DeployContractWithConfFilePath(configFilePath string) {
	conf := config.LoadConfig(configFilePath)

	// 1. 创建节点
	node := NewNode(
		conf.Nodes[0].Addr,
		conf.Nodes[0].ConnCnt,
		conf.Nodes[0].UseTls,
		conf.Nodes[0].CaPaths,
		conf.Nodes[0].TlsHostName,
	)

	// 2. 创建客户端
	client, err := NewClient(
		node,
		conf.Client.ChainClientOrgId,
		conf.Client.ChainId,
		conf.Client.AuthType,
		conf.Client.UserKeyFilePath,
		conf.Client.UserCrtFilePath,
	)
	if err != nil {
		panic(err)
	}

	// 3. 读取abi
	myAbi, err := PrepareAbi(conf.Contract.AbiPath)
	if err != nil {
		panic(err)
	}

	// 4. 根据合约构造函数是否需要参数，进行Pack
	packData, err := Pack(myAbi, "")
	if err != nil {
		panic(err)
	}

	// 5. 读取bin
	binCode, err := PrepareBin(conf.Contract.BinPath)
	if err != nil {
		panic(err)
	}

	// 6. 创建payload
	payload, err := CreatePayload(
		client,
		conf.Contract.Name,
		conf.Contract.Version,
		string(binCode),
		packData,
		common.RuntimeType_EVM,
	)

	// 7. 创建背书，进行多签
	endorsers, err := Endor(client, payload, conf.Endorsers)
	if err != nil {
		panic(err)
	}

	// 8. 部署合约
	resp, err := DeployContract(client, payload, endorsers, 5, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	// 9. 检查回执
}

// NewNode 创建一个节点.
func NewNode(addr string, connCnt int, userTls bool, caPaths []string, tlsHostName string) *sdk.NodeConfig {
	node := sdk.NewNodeConfig(
		sdk.WithNodeAddr(addr),
		sdk.WithNodeConnCnt(connCnt),
		sdk.WithNodeUseTLS(userTls),
		sdk.WithNodeCAPaths(caPaths),
		sdk.WithNodeTLSHostName(tlsHostName),
	)

	return node
}

// NewClient 根据节点创建一个客户端.
func NewClient(node *sdk.NodeConfig, orgId, chainId, authType, userKeyFilePath, userCrtFilePath string) (*sdk.ChainClient, error) {
	chainClient, err := sdk.NewChainClient(
		sdk.WithChainClientOrgId(orgId),
		sdk.WithChainClientChainId(chainId),
		sdk.WithAuthType(authType),
		sdk.WithUserKeyFilePath(userKeyFilePath),
		sdk.WithUserCrtFilePath(userCrtFilePath),
		sdk.AddChainClientNodeConfig(node),
	)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	err = chainClient.EnableCertHash()
	if err != nil {
		log.Fatal(err)
	}

	return chainClient, nil
}

// PrepareAbi 读取abi, 构造abi对象.
func PrepareAbi(abiPath string) (abi.ABI, error) {
	abiJson, err := ioutil.ReadFile(abiPath)
	if err != nil {
		log.Println(err)

		return abi.ABI{}, err
	}

	myAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		log.Println(err)

		return abi.ABI{}, err
	}

	return myAbi, nil
}

// Pack.
func Pack(myAbi abi.ABI, name string, args ...interface{}) ([]byte, error) {
	dataByte, err := myAbi.Pack(name, args...) // FIXME
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return dataByte, nil
}

// 5. 读取bin.
func PrepareBin(binPath string) ([]byte, error) {
	return ioutil.ReadFile(binPath)
}

// CalcContractName 将合约名进行hash.
func CalcContractName(contractName string) string {
	return hex.EncodeToString(evmutils.Keccak256([]byte(contractName)))[24:]
}

// 6. 创建Payload.
func CreatePayload(client *sdk.ChainClient, contractName, version, binCode string, packByte []byte, runtimeType common.RuntimeType) (*common.Payload, error) {
	contractName = CalcContractName(contractName)
	data := hex.EncodeToString(packByte)
	pairs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(data),
		},
	}

	payload, err := client.CreateContractCreatePayload(
		contractName,
		version,
		binCode,
		common.RuntimeType_EVM,
		pairs)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return payload, nil
}

// 7. 创建背书，进行多签, 只支持PermissionedWithCert.
func Endor(client *sdk.ChainClient,
	payload *common.Payload,
	endorserUsers []*config.User,
) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	for _, user := range endorserUsers {
		var entry *common.EndorsementEntry
		var err error
		entry, err = sdkutils.MakeEndorserWithPath(user.SignKeyPath, user.SignCrtPath, payload)
		if err != nil {
			return nil, err
		}

		endorsers = append(endorsers, entry)
	}

	return endorsers, nil
}

// 8. 部署合约.
func DeployContract(client *sdk.ChainClient, payload *common.Payload, endorsers []*common.EndorsementEntry, createContractTimeout int64, withSyncResult bool) (*common.TxResponse, error) {
	// 发送创建合约请求
	resp, err := client.SendContractManageRequest(payload, endorsers, createContractTimeout, withSyncResult)
	if err != nil {
		return nil, err
	}

	err = CheckProposalRequestResp(resp, true)
	if err != nil {
		return nil, err
	}
	fmt.Println("---")

	return resp, nil
}

func CheckProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		if resp.Message == "" {
			resp.Message = resp.Code.String()
		}
		return errors.New(resp.Message)
	}

	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}

	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return errors.New(resp.ContractResult.Message)
	}

	return nil
}
