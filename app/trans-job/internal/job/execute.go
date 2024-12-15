package job

import (
	"context"
	"encoding/json"

	"chainmaker.org/chainmaker/pb-go/v2/common"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/pkg/task"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/gotools"
)

func (j *TransJob) Name() string {
	return "vtrace-trans-job"
}

/*
TODO

1. 成功但未确认的消息处理:处理前需要查询数据库，将status为success的直接确认
2. 处理多次但均未成功的消息:强行执行.
*/
func (j *TransJob) ExecuteJob(ctx context.Context) error {
	jobs, err := j.mq.ReadGroup(ctx, constants.MqStream, constants.MqGroup, constants.MqConsumer, 10, 0)
	if err != nil {
		logger.Error(err)

		return err
	}

	if jobs == nil || len(jobs) <= 0 {
		logger.Info("no job found")

		return nil
	}

	for taskId, job := range jobs {
		gotools.Go(func(taskId string, job []byte) func() {
			return func() {
				// 0. 检查任务
				if len(job) <= 0 {
					return
				}
				logger.Infof("execute task, Id: %s, Value: %s", taskId, job)
				var t task.TransTask
				if err := json.Unmarshal(job, &t); err != nil {
					logger.Error(err)

					return
				}

				// 1. 创建client
				client, err := sdk.NewChainClient(
					sdk.WithChainClientOrgId(t.OrgId),
					sdk.WithChainClientChainId(t.ChainId),
					sdk.WithAuthType(j.AuthType),
					sdk.WithUserKeyBytes(t.Pk),
					sdk.WithUserCrtBytes(t.Crt),
					sdk.WithUserSignCrtBytes(t.Crt),
					sdk.WithUserSignKeyBytes(t.Pk),
					sdk.AddChainClientNodeConfig(j.node[0]),
				)
				if err != nil {
					logger.Errorf("error create client: %+v", err)

					return
				}

				// 2. 发送交易
				resp, err := client.InvokeContract(
					t.ContractName,
					t.Method,
					"",
					t.Kv,
					-1,
					true,
				)
				if err != nil {
					logger.Errorf("invoke contract failed: %+v", err)

					return
				}

				success := resp.Code == common.TxStatusCode_SUCCESS

				if !success {
					logger.Errorf("链上交易失败：%v\n", resp.ContractResult.Message)
				}

				// 3. 获取交易数据
				res, err := j.ContractData.Abi.Unpack(t.MethodName, resp.ContractResult.Result)
				if err != nil {
					logger.Errorf("failed unpack resp: %+v", err)

					return
				}

				logger.Infof("trans result: %+v, txId: %s", res, resp.TxId)

				// 4. 更新trans记录
				if _, err := j.Trans().UpdateTrans(ctx, &transV1.UpdateTransRequest{
					TransId: t.TaskId,
					TxHash:  resp.TxId,
					Success: success,
				}); err != nil {
					logger.Error(err)

					return
				}

				// 5. 更新circ记录
				var status circV1.RecordStatus
				if success {
					status = circV1.RecordStatus_success
				} else {
					status = circV1.RecordStatus_failed
				}

				if _, err := j.Circ().UpdateCircStatus(ctx, &circV1.UpdateCircStatusRequest{
					TransId: t.TaskId,
					Status:  status,
				}); err != nil {
					logger.Error(err)

					return
				}

				if success {
					j.cache.Set(ctx, t.TaskId, resp.GetTxId(), &cache.Options{})
				}

				if err := j.mq.Ack(ctx, constants.MqStream, constants.MqGroup, taskId); err != nil {
					logger.Error(err)

					return
				}
			}
		}(taskId, job))
	}

	return nil
}
