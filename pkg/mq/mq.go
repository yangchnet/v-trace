package mq

import "context"

// Mq 消息队列.
//
//go:generate mockgen -source=mq.go -destination=mock_mq.go -package=mq . Mq
type Mq interface {
	Publisher
	Consumer
}

// Publisher add msg to mq.
type Publisher interface {
	// Add 向mq队列中发送一个消息id为msgId的[key:value]消息
	Add(ctx context.Context, mq, msgId, key string, value []byte) error
}

// Consumer consume msg from mq and ack it.
type Consumer interface {
	// CreateGroup 创建一个消费组
	CreateGroup(ctx context.Context, mq, group string) error

	// ReadGroup 以group中consumer的身份读取mq中从start开始的count个信息
	ReadGroup(ctx context.Context, mq, group, consumer string, count int64, start int) (map[string][]byte, error)

	// DestroyGroup
	DestroyGroup(ctx context.Context, mq, group string) error

	// Pending 查询mq中group组中从start到end的count个信息
	// Pending(ctx context.Context, mq, group string, start, end int, count int64) ([][]byte, error)

	// Ack 向mq队列发送确认信息，告知其group组中消息键为key的消息已被成功消费
	Ack(ctx context.Context, mq, group string, key string) error
}
