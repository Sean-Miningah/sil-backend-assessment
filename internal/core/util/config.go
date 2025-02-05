package util

import (
	"github.com/spf13/viper"
)

const (
	EnvProd = "prod"
	EnvDev = "dev"
)

type Config struct {
	Env string `mapstructure:"ENV"`

	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`

	ServerType string `mapstructure:"SERVER_TYPE"`
	Port int `mapstructure:"PORT"`

	TokenSecret string `mapstructure:"TOKEN_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}