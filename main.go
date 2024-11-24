package main

import (
	"context"
	"fmt"
	"log" // to validate urls
	"os"  // env

	// to create time variables
	"github.com/joho/godotenv" // env

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"

	// "github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

// TableBasics encapsulates the Amazon DynamoDB service actions.
// Contains a DynamoDB service client that is used to act on the specified table.
type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

var logger *zap.Logger
var basics TableBasics

const (
	HASH_LENGTH = 6
)

func init() {
	var err error

	// Load environment variables from .env
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Our logger
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}

	// Read AWS credentials and region from environment
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

	if awsAccessKey == "" || awsSecretKey == "" || awsRegion == "" {
		log.Fatalf("AWS credentials or region not set in .env file")
	}

	// Load AWS configuration with credentials provider
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		logger.Fatal("Failed to load AWS config", zap.Error(err))
	}

	// Our DynamoDB session
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// Set up TableBasics with the DynamoDB client and table name
	basics = TableBasics{
		DynamoDbClient: dynamoClient,
		TableName:      "ShortUrls",
	}
}

func main() {
	app := fiber.New()
	fmt.Println("Server started")

	urlsRoutes := app.Group("/urls")

	// Routes
	urlsRoutes.Post("/shorten", shortenURL)
	urlsRoutes.Get("/:shortCode", getURL)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
