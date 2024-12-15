package mq

import (
	"context"
	"fmt"

	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type RedisMqConfig struct {
	Host string `mapstructure:"host"`
	Port int64  `mapstructure:"port"`
}

// RedisMq 使用redis Stream作为消息队列.
type RedisMq struct {
	client *redis.Client
}

func NewRedisMq(ctx context.Context, redisMqConfig *RedisMqConfig) *RedisMq {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisMqConfig.Host, redisMqConfig.Port),
	})

	return &RedisMq{
		client: client,
	}
}

var _ Mq = (*RedisMq)(nil)

// Add 向mq队列中发送一个消息id为msgId的[key:value]消息.
func (q *RedisMq) Add(ctx context.Context, mq string, msgId string, key string, value []byte) error {
	if err := q.client.XAdd(ctx, &redis.XAddArgs{
		Stream: mq,
		Values: map[string]interface{}{
			key: value,
		},
	}).Err(); err != nil {
		logger.Error(err)

		return err
	}

	return nil
}

// CreateGroup 创建一个消费组.
func (q *RedisMq) CreateGroup(ctx context.Context, mq, group string) error {
	if err := q.client.XGroupCreateMkStream(ctx, mq, group, "0").Err(); err != nil {
		if err.Error() == "BUSYGROUP Consumer Group name already exists" {
			return nil
		}
		logger.Error(err)

		return err
	}

	return nil
}

// ReadGroup 以group中consumer的身份读取mq中从start开始的count个信息.
func (q *RedisMq) ReadGroup(ctx context.Context, mq string, group string, consumer string, count int64, start int) (map[string][]byte, error) {
	entries, err := q.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams: []string{
			mq,
			">",
		},
		Count: count,
		Block: 0,
		NoAck: false,
	}).Result()
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	res := make(map[string][]byte)
	for i := 0; i < len(entries[0].Messages); i++ {
		messageID := entries[0].Messages[i].ID
		msgValue := entries[0].Messages[i].Values
		value := msgValue[constants.MqKey].(string)

		res[messageID] = []byte(value)
	}

	return res, nil
}

func (q *RedisMq) DestroyGroup(ctx context.Context, mq, group string) error {
	if err := q.client.XGroupDestroy(ctx, mq, group).Err(); err != nil {
		logger.Error(err)

		return err
	}

	return nil
}

// Pending 查询mq中group组中从start到end的count个信息.
// func (q *RedisMq) Pending(ctx context.Context, mq string, group string, start int, end int, count int64) ([][]byte, error) {
// 	panic("not implemented") // TODO: Implement
// }

// Ack 向mq队列发送确认信息，告知其group组中消息键为key的消息已被成功消费.
func (q *RedisMq) Ack(ctx context.Context, mq string, group string, ids string) error {
	result, err := q.client.XAck(ctx, mq, group, ids).Result()
	if err != nil {
		logger.Error(err)

		return err
	}

	logger.Infof("mq ack result: %d", result)

	// if err := q.client.XAck(ctx, mq, group, ids).Err(); err != nil {
	// 	logger.Error(err)

	// 	return err
	// }

	return nil
}
