package contract

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func NewNodes(nodeConfig []*config.NodeConfig) []*sdk.NodeConfig {
	if len(nodeConfig) <= 0 {
		panic("node config not set")
	}

	nodes := make([]*sdk.NodeConfig, 0)
	for _, nc := range nodeConfig {
		nodes = append(nodes, sdk.NewNodeConfig(
			sdk.WithNodeAddr(nc.Addr),
			sdk.WithNodeConnCnt(nc.ConnCnt),
			sdk.WithNodeUseTLS(nc.UseTls),
			sdk.WithNodeCAPaths(nc.CaPaths),
			sdk.WithNodeTLSHostName(nc.TlsHostName),
		))
	}

	return nodes
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

func NewChainClient(node *sdk.NodeConfig, orgId, chainId, authType string, user *config.User) (*sdk.ChainClient, error) {
	chainClient, err := sdk.NewChainClient(
		sdk.WithChainClientOrgId(orgId),
		sdk.WithChainClientChainId(chainId),
		sdk.WithAuthType(authType),
		sdk.WithUserKeyFilePath(user.TlsKeyPath),
		sdk.WithUserCrtFilePath(user.TlsCrtPath),
		sdk.WithUserSignCrtFilePath(user.SignCrtPath),
		sdk.WithUserSignKeyFilePath(user.SignKeyPath),
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
