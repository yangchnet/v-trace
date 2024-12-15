package app

type AppConfig struct {
	ID        string   `mapstructure:"id"`
	Name      string   `mapstructure:"name"`
	Version   string   `mapstructure:"version"`
	Endpoints []string `mapstructure:"endpoints"`
}
