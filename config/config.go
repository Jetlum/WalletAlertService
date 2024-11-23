package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL     string
	SendGridAPIKey  string
	InfuraProjectID string
}

func LoadConfig() (*Config, error) {
	if os.Getenv("GO_ENV") == "test" {
		return &Config{
			DatabaseURL:     "test_db_url",
			SendGridAPIKey:  "test_api_key",
			InfuraProjectID: "test_infura_id",
		}, nil
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
