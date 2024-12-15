package config

import "github.com/spf13/viper"

type Config struct {
	Nodes          []*NodeConfig  `mapstructure:"nodes"`
	Client         ClientConfig   `mapstructure:"client"`
	Contract       ContractConfig `mapstructure:"contract"`
	Endorsers      []*User        `mapstructure:"endorsers"`
	WithSyncResult bool           `mapstructure:"with_sync_result"`
}

type NodeConfig struct {
	Addr        string   `mapstructure:"addr"`
	ConnCnt     int      `mapstructure:"conn_cnt"`
	UseTls      bool     `mapstructure:"use_tls"`
	CaPaths     []string `mapstructure:"ca_paths"`
	TlsHostName string   `mapstructure:"tls_hostname"`
}

type ClientConfig struct {
	ChainClientOrgId string `mapstructure:"chain_client_org_id"`
	ChainId          string `mapstructure:"chain_id"`
	AuthType         string `mapstructure:"auth_type"`
	UserKeyFilePath  string `mapstructure:"user_key_file_path"`
	UserCrtFilePath  string `mapstructure:"user_crt_file_path"`
}

type ContractConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	AbiPath string `mapstructure:"abi_path"`
	BinPath string `mapstructure:"bin_path"`
}

type User struct {
	SignKeyPath string `mapstructure:"user_key"`
	SignCrtPath string `mapstructure:"user_crt"`
	TlsKeyPath  string `mapstructure:"user_tls_key"`
	TlsCrtPath  string `mapstructure:"user_tls_crt"`
}

func LoadConfig(configFilePath string) *Config {
	var conf Config = Config{}
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.Unmarshal(&conf)

	return &conf
}
