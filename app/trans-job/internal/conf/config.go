package conf

import (
	cmConfig "gitee.com/qciip-icp/v-trace/app/pkg/contract/config"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/cron"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/mq"
	"gitee.com/qciip-icp/v-trace/pkg/pubsub"
	"gitee.com/qciip-icp/v-trace/pkg/registry"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	Cron       cron.CronConfig          `mapstructure:"cron"`
	Pubsub     pubsub.RedisPubsubConfig `mapstructure:"pubsub"`
	Mq         mq.RedisMqConfig         `mapstructure:"mq"`
	Cache      cache.RedisConfig        `mapstructure:"redis"`
	ChainMaker cmConfig.Config          `mapstructure:"chainmaker"`
	Log        logger.Config            `mapstructure:"log"`
	Etcd       registry.Etcd            `mapstructure:"etcd"`
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
