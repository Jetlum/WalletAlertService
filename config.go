package config

type Config struct {
	InfuraProjectID string
	DatabaseURL     string
	SendGridAPIKey  string
}

func LoadConfig() Config {
	return Config{
		InfuraProjectID: "YOUR_INFURA_PROJECT_ID",
		DatabaseURL:     "DatabaseURL",
		SendGridAPIKey:  "SendGridAPIKey",
	}
}
