package config

import (
	// We cant use zap.Logger as this is the second file to be called on start
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel     string `envconfig:"LOG_LEVEL"`
	AwsAccessKey string `envconfig:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion    string `envconfig:"AWS_REGION"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(
		func() {
			// Load environment variables from .env
			err := envconfig.Process("", &config)
			if err != nil {
				log.Fatalf("Error loading .env file: %v", err)
			}
			configBytes, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				log.Fatalf("AWS credentials or region not set in .env file")
			}
			fmt.Println("Configuration:", string(configBytes))
		},
	)
	return &config
}
