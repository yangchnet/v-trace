package biz

import (
	"context"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/conf"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewTransCase, NewNode)

type TransCase struct {
	Nodes []*sdk.NodeConfig

	ContractData *contract.ContractData

	AuthType string

	repo TransRepo
}

func NewNode(ctx context.Context, c *conf.Bootstrap) []*sdk.NodeConfig {
	nodes := make([]*sdk.NodeConfig, 0)
	for _, node := range c.ChainMaker.Nodes {
		nodes = append(nodes, contract.NewNode(
			node.Addr,
			node.ConnCnt,
			node.UseTls,
			node.CaPaths,
			node.TlsHostName))
	}

	return nodes
}

func NewTransCase(repo TransRepo, config *conf.Bootstrap, nodes ...*sdk.NodeConfig) *TransCase {
	contractData, err := contract.NewContractData(
		config.ChainMaker.Contract.AbiPath,
		config.ChainMaker.Contract.BinPath,
		config.ChainMaker.Contract.Name,
	)
	if err != nil {
		panic(err)
	}

	if len(nodes) <= 0 {
		panic("the number of nodes cannot be 0")
	}

	return &TransCase{
		Nodes:        nodes,
		ContractData: contractData,
		repo:         repo,
		AuthType:     config.ChainMaker.Client.AuthType,
	}
}
