package config

import "os"

type Config struct {
	APIKey string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("API_KEY")

	config := &Config{
		APIKey: apiKey,
	}

	return config, nil
}
