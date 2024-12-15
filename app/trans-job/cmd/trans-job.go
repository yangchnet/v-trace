package main

import (
	"context"
	"time"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract"
	"gitee.com/qciip-icp/v-trace/app/trans-job/internal/conf"
	"gitee.com/qciip-icp/v-trace/pkg/cron"
	"gitee.com/qciip-icp/v-trace/pkg/mq"
)

func NewNodes(ctx context.Context, c *conf.Bootstrap) []*sdk.NodeConfig {
	return contract.NewNodes(c.ChainMaker.Nodes)
}

func NewConsumer(ctx context.Context, c *conf.Bootstrap) mq.Consumer {
	return mq.NewRedisMq(ctx, &c.Mq)
}

func main() {
	time.Sleep(10 * time.Second) // sleep 10s 等待stream创建完成

	ctx := context.Background()

	c := conf.LoadConfig("app/trans-job/config")

	transJob, err := InitTransJob(ctx, c)
	if err != nil {
		panic(err)
	}

	cr := cron.NewCron(c.Cron.CronName)

	cr.Register(transJob.Name(), transJob)

	for jobName, spec := range c.Cron.Jobs {
		cr.StartJob(ctx, spec, jobName)
	}
	select {}
}
