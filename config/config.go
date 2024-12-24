package config

import (
	// We cant use zap.Logger as this is the second file to be called on start
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel           string `envconfig:"LOG_LEVEL"`
	ServerReadTimeout  string `envconfig:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout string `envconfig:"SERVER_WRITE_TIMEOUT"`
	AwsAccessKey       string `envconfig:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey       string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion          string `envconfig:"AWS_REGION"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(
		func() {
			// Load environment variables from .env file
			err := godotenv.Load(".env") // Specify the correct path to your .env file
			if err != nil {
				log.Fatalf("Error loading .env file: %v", err)
			}

			// Load environment variables from .env
			err = envconfig.Process("", &config)
			if err != nil {
				log.Fatalf("Error loading env variables from .env or exported env variables: %v", err)
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

func GetFiberConfig() fiber.Config {
	appConfig := Get()

	timeToDur := func(configValue string) time.Duration {
		intValue, _ := strconv.Atoi(configValue)
		return time.Second * time.Duration(intValue)
	}

	return fiber.Config{
		ReadTimeout:  timeToDur(appConfig.ServerReadTimeout),
		WriteTimeout: timeToDur(appConfig.ServerWriteTimeout),
	}
}
