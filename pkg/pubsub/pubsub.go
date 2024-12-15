package pubsub

import "context"

// Pubsub 实现了发布订阅机制.
//
//go:generate mockgen -source=pubsub.go -destination=mock_pubsub.go -package=pubsub . Pubsub
type Pubsub interface {
	// Publish 发布
	Publish(ctx context.Context, chName string, payload []byte) error
	// Subscribe 订阅
	Subscribe(ctx context.Context, chName string) error
	// Receive 每次调用Receive将从chan中获取一个信息.
	Receive(ctx context.Context) ([][]byte, error)
	// ChName 获取频道名称
	ChName() string
}

// UnimplementedPubsub is a unimplement pubsub.
type UnimplementedPubsub struct{}

// Publish 发布.
func (p *UnimplementedPubsub) Publish(ctx context.Context, chName string, payload []byte) error {
	panic("not implemented") // TODO: Implement
}

// Subscribe 订阅.
func (p *UnimplementedPubsub) Subscribe(ctx context.Context, chName string) error {
	panic("not implemented") // TODO: Implement
}

// Receive 每次调用Receive将从chan中获取一个信息.
func (p *UnimplementedPubsub) Receive(ctx context.Context) ([][]byte, error) {
	panic("not implemented") // TODO: Implement
}

// ChName 获取频道名称.
func (p *UnimplementedPubsub) ChName() string {
	panic("not implemented") // TODO: Implement
}
