package conf

import (
	"gitee.com/qciip-icp/v-trace/app/pkg/algo"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	Grpc grpc.GrpcConfig `mapstructure:"grpc"`
	Algo algo.AlgoConfig `mapstructure:"model"`
	Data DataConfig      `mapstructure:"data"`
	Log  logger.Config   `mapstructure:"log"`
	Etcd registry.Etcd   `mapstructure:"etcd"`
}

type DataConfig struct {
	Db    DBConfig          `mapstructure:"database"`
	Redis cache.RedisConfig `mapstructure:"redis"`
}

type DBConfig struct {
	Driver string `mapstructure:"driver"`
	Dsn    string `mapstructure:"dsn"`
}

func LoadConfig(configFilePath string) *Bootstrap {
	var conf Bootstrap = Bootstrap{}
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	logger.SetLogger(&conf.Log, nil)

	return &conf
}
