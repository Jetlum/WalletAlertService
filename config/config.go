package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL     string
	SendGridAPIKey  string
	InfuraProjectID string
	CoinGeckoAPIKey string
}

func LoadConfig() (*Config, error) {
	if os.Getenv("GO_ENV") == "test" {
		return &Config{
			DatabaseURL:     "test_db_url",
			SendGridAPIKey:  "test_api_key",
			InfuraProjectID: "test_infura_id",
			CoinGeckoAPIKey: "test_coingecko_key",
		}, nil
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		DatabaseURL:     viper.GetString("database.url"),
		SendGridAPIKey:  viper.GetString("sendgrid.api_key"),
		InfuraProjectID: viper.GetString("infura.project_id"),
		CoinGeckoAPIKey: viper.GetString("coingecko.api_key"),
	}

	return &cfg, nil
}
