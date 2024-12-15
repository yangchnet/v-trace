package mq

import (
	"context"
	"testing"

	"gitee.com/qciip-icp/v-trace/pkg/constants"
)

func Test_Mq(t *testing.T) {
	ctx := context.Background()
	q := NewRedisMq(ctx, &RedisMqConfig{
		Host: "vtrace_redis",
		Port: 16379,
	})

	// if err := q.Add(ctx, "s1", "", "k1", []byte("111111111111")); err != nil {
	// 	t.Fatal(err)
	// }

	// if err := q.Add(ctx, "s1", "", "k1", []byte("22222222222222")); err != nil {
	// 	t.Fatal(err)
	// }

	// if err := q.Add(ctx, "s1", "", "k1", []byte("333333333333")); err != nil {
	// 	t.Fatal(err)
	// }

	if err := q.CreateGroup(ctx, constants.MqStream, constants.MqGroup); err != nil {
		t.Fatal(err)
	}

	// res, err := q.ReadGroup(ctx, "s1", "g1", "c1", 10, 0)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// t.Log(res)

	// for k, _ := range res {
	// 	if err := q.Ack(ctx, "s1", "g1", k); err != nil {
	// 		t.Fatal(err)
	// 	}
	// }
}
