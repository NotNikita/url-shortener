package store

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	appConfig "url-shortener/config"
	"url-shortener/logger"
	awsStore "url-shortener/store/aws"
)

// Store encapsulates the Amazon DynamoDB service actions.
// Contains a DynamoDB service client that is used to act on the specified table.
type Store struct {
	DynamoDbClient *dynamodb.Client

	Urls UrlsRepo
}

func NewStore(ctx *fiber.Ctx) (*Store, error) {
	appCfg := appConfig.Get()
	logger := logger.Get()
	// Load AWS configuration with credentials provider
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(appCfg.AwsRegion),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(appCfg.AwsAccessKey, appCfg.AwsSecretKey, "")),
	)
	if err != nil {
		logger.Fatal("Failed to load AWS config", zap.Error(err))
	}

	// Our DynamoDB session
	dynamoClient := dynamodb.NewFromConfig(awsCfg)

	return &Store{
		DynamoDbClient: dynamoClient,
		Urls:           awsStore.NewUrlsRepo(dynamoClient),
	}, nil
}
