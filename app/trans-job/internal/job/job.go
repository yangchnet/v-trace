package job

import (
	"context"
	"fmt"

	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	localGrpc "gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/registry"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/app/trans-job/internal/conf"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/mq"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewTransJob, NewCache)

type TransJob struct {
	cache cache.Cache
	node  []*sdk.NodeConfig

	mq mq.Consumer

	ContractData *contract.ContractData

	AuthType string

	connGetter map[string]localGrpc.ConnGetInterface
}

func NewCache(ctx context.Context, c *conf.Bootstrap) *cache.RedisStore {
	return cache.NewRedisStore(&c.Cache)
}

func NewTransJob(
	ctx context.Context,
	config *conf.Bootstrap,
	mq mq.Consumer,
	cache cache.Cache,
	r registry.Registrar,
	node ...*sdk.NodeConfig,
) *TransJob {
	contractData, err := contract.NewContractData(
		config.ChainMaker.Contract.AbiPath,
		config.ChainMaker.Contract.BinPath,
		config.ChainMaker.Contract.Name,
	)
	if err != nil {
		panic(fmt.Sprintf("prepare contract data filed: %+v", err))
	}

	// 循环5次创建group
	success := false
	for i := 0; i < 5; i++ {
		if err := mq.CreateGroup(ctx, constants.MqStream, constants.MqGroup); err != nil {
			logger.Errorf("create group failed: %+v, group name: %s", err, constants.MqGroup)

			continue
		}
		success = true

		break
	}

	if !success {
		panic("error create group")
	}

	logger.Infof("current mq group: %s", constants.MqGroup)

	getter := make(map[string]grpc.ConnGetInterface)
	transWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Trans)
	if err != nil {
		panic(err)
	}
	getter[constants.Trans] = localGrpc.NewConnGetter(ctx, transWatcher, constants.Trans)

	circWatcher, err := r.(*registry.Registry).Watch(ctx, constants.Circ)
	if err != nil {
		panic(err)
	}
	getter[constants.Circ] = localGrpc.NewConnGetter(ctx, circWatcher, constants.Circ)

	return &TransJob{
		cache:        cache,
		node:         node,
		ContractData: contractData,
		AuthType:     constants.ContractAuthType,
		mq:           mq,
		connGetter:   getter,
	}
}

func (j *TransJob) Trans() transV1.TransServiceClient {
	conn := j.connGetter[constants.Trans].Get()

	return transV1.NewTransServiceClient(conn)
}

func (j *TransJob) Circ() circV1.CircServiceClient {
	conn := j.connGetter[constants.Circ].Get()

	return circV1.NewCircServiceClient(conn)
}
