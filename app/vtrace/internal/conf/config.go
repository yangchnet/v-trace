package conf

import (
	"time"

	"gitee.com/qciip-icp/v-trace/pkg/fs/qiniuoss"

	"gitee.com/qciip-icp/v-trace/pkg/grpc"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"gitee.com/qciip-icp/v-trace/pkg/token"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	Gateway GatewayConfig           `mapstructure:"gateway"`
	Grpc    grpc.GrpcConfig         `mapstructure:"grpc"`
	Log     logger.Config           `mapstructure:"log"`
	Etcd    registry.Etcd           `mapstructure:"etcd"`
	OSS     qiniuoss.QiniuOSSConfig `mapstructure:"qiniuoss"`
}

type GatewayConfig struct {
	Http  HttpConfig        `mapstructure:"http"`
	Log   logger.Config     `mapstructure:"log"`
	Token token.TokenConfig `mapstructure:"token"`
}

type HttpConfig struct {
	Host    string        `mapstructure:"host"`
	Port    int64         `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
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
