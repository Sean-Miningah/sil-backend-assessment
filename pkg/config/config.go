package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	Address        string `mapstructure:"SERVER_ADDRESS"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	ServiceName    string `mapstructure:"TELEMETRY_SERVICE_NAME"`
	JaegerEndpoint string `mapstructure:"JAEGER_ENDPOINT"`
	PrometheusPort string `mapstructure:"PROMETHEUS_PORT"`

	// Zitadel fields
	// ZitadelProjectID  string `mapstructure:"ZITADEL_PROJECT_ID"`
	ZitadelIssuerURL    string `mapstructure:"ZITADEL_ISSUER_URL"`
	ZitadelClientID     string `mapstructure:"ZITADEL_CLIENT_ID"`
	ZitadelClientSecret string `mapstructure:"ZITADEL_CLIENT_SECRET"`
	ZitadelRedirectURI  string `mapstructure:"ZITADEL_REDIRECT_URI"`
	// CookieSecret      string `mapstructure:"COOKIE_SECRET"`
}

func Load(path string) *Config {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)

	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &config
}
