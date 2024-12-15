package algo

import (
	"context"

	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/google/wire"
	"github.com/yangchnet/tf-serving/tensorflow_serving/apis"
	"google.golang.org/grpc"
)

var AlgoProviderSet = wire.NewSet(NewAlgoClient)

type AlgoConfig struct {
	Address string `mapstructure:"address"`
}

type AlgoClient struct {
	apis.ModelServiceClient
	apis.PredictionServiceClient
	apis.SessionServiceClient
}

func NewAlgoClient(ctx context.Context, config *AlgoConfig) *AlgoClient {
	conn, err := grpc.DialContext(ctx, config.Address, grpc.WithInsecure())
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	modelClient := apis.NewModelServiceClient(conn)
	predictClient := apis.NewPredictionServiceClient(conn)
	sessionClient := apis.NewSessionServiceClient(conn)

	return &AlgoClient{
		modelClient,
		predictClient,
		sessionClient,
	}
}
