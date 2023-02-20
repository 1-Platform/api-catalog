package config

import "github.com/spf13/viper"

type Config struct {
	MongoDbURI string `mapstructure:"MONGO_DB_URI"`
}

func New(path string) (*Config, error) {
	var config Config
	// load config paths
	viper.AddConfigPath("~/")
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")
	viper.SetConfigName("apic")

	// load from env if found
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
