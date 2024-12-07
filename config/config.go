package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL        string
	SendGridAPIKey     string
	InfuraProjectID    string
	CoinGeckoAPIKey    string // Add CoinGecko API key
	PriceCheckInterval int    // Interval in minutes for price checking
}

func LoadConfig() (*Config, error) {
	if os.Getenv("GO_ENV") == "test" {
		return &Config{
			DatabaseURL:     os.Getenv("TEST_DATABASE_URL"),
			SendGridAPIKey:  os.Getenv("TEST_SENDGRID_API_KEY"),
			InfuraProjectID: os.Getenv("TEST_INFURA_PROJECT_ID"),
			CoinGeckoAPIKey: os.Getenv("TEST_COINGECKO_API_KEY"),
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
