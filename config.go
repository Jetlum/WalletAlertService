package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	InfuraProjectID string `mapstructure:"infura.project_id"`
	DatabaseURL     string `mapstructure:"database.url"`
	SendGridAPIKey  string `mapstructure:"sendgrid.api_key"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		InfuraProjectID: viper.GetString("infura.project_id"),
		DatabaseURL:     viper.GetString("database.url"),
		SendGridAPIKey:  viper.GetString("sendgrid.api_key"),
	}, nil
}
