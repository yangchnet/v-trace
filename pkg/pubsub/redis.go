package pubsub

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

var RedisMqProvider = wire.NewSet(NewRedisPubsub)

type RedisPubsubConfig struct {
	Host   string `mapstructure:"host"`
	Port   int64  `mapstructure:"port"`
	ChName string `mapstructure:"ch_name"`
}

// RedisPubsub.
type RedisPubsub struct {
	client *redis.Client
	chName string
	pubsub *redis.PubSub
}

// NewRedisPubsub 创建一个新的发布/订阅.
func NewRedisPubsub(ctx context.Context, psConfig *RedisPubsubConfig, subfunc func(ps Pubsub)) Pubsub {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", psConfig.Host, psConfig.Port),
	})

	ps := &RedisPubsub{
		client: client,
		chName: psConfig.ChName,
	}

	if subfunc != nil {
		subfunc(ps)
	}

	return ps
}

// Publish 发布.
func (q *RedisPubsub) Publish(ctx context.Context, chName string, payload []byte) error {
	return q.client.Publish(ctx, chName, payload).Err()
}

// Subscribe 订阅.
func (q *RedisPubsub) Subscribe(ctx context.Context, chName string) error {
	pubsub := q.client.Subscribe(ctx, chName)
	q.pubsub = pubsub

	return nil
}

// Receive 每次调用Receive将从mq中获取10个信息.
// 调用Receive之前必须调用Subscribe.
func (q *RedisPubsub) Receive(ctx context.Context) ([][]byte, error) {
	msgCh := q.pubsub.Channel()

	res := make([][]byte, 0)
	for i := 0; i < len(msgCh); i++ {
		if i >= 10 {
			break
		}
		msg := <-msgCh

		res = append(res, []byte(msg.Payload))
	}
	return res, nil
}

// ChName 获取频道名称.
func (q *RedisPubsub) ChName() string {
	return q.chName
}
