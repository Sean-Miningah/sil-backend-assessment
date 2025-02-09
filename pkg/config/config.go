package config

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Server      ServerConfig
	Database    DatabaseConfig
	Telemetry   TelemetryConfig
}

type ServerConfig struct {
	Address string `mapstructure:"SERVER_ADDRESS"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type TelemetryConfig struct {
	ServiceName    string `mapstructure:"TELEMETRY_SERVICE_NAME"`
	JaegerEndpoint string `mapstructure:"JAEGER_ENDPOINT"`
	PrometheusPort string `mapstructure:"PROMETHEUS_PORT"`
}

func Load() *Config {
	viper.SetConfigFile(".env.dev")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
