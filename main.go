package main

import (
	"fmt"
	"log"
	"net/url" // to validate urls
	"time"    // to create time variables
	"context"
	"github.com/joho/godotenv"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go/aws/awserr"
	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofiber/fiber/v2"
	// "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// TableBasics encapsulates the Amazon DynamoDB service actions used in the examples.
// It contains a DynamoDB service client that is used to act on the specified table.
type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

// Movie encapsulates data about a movie. Title and Year are the composite primary key
// of the movie in Amazon DynamoDB. Title is the sort key, Year is the partition key,
// and Info is additional data.
type UrlData struct {
	ShortUrl    string `dynamodbav:"shortCode"`
	OriginalUrl string `dynamodbav:"originalURL"`
	CreatedAt   string `dynamodbav:"createdAt"`
	ExpiresAt   string `dynamodbav:"expiresAt"`
}

var logger *zap.Logger
var basics TableBasics

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

	// Routes
	app.Post("/shorten", shortenURL)
	app.Get("/:shortCode", getURL)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func isValidUrl(urlToValidate string) bool {
	u, err := url.Parse(urlToValidate)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func shortenURL(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Fatalf("Error parsing provided data for shortenURL: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form: " + err.Error(),
		})
	}

	long_url := form.Value["long_url"]
	if len(long_url)== 0 || !isValidUrl(long_url[0]) {
		log.Fatalf("Error when validating provided url to shorten")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to validate url: " + err.Error(),
		})
	}

	now := time.Now().Format(time.RFC3339)
	nowPlusMonth := time.Now().AddDate(0,1,0).Format(time.RFC3339)
	// put valid in db
	// TODO: short data xxHash creation
	urlData := UrlData{
		ShortUrl:    "short-" + time.Now().Format(time.RFC3339),
		OriginalUrl: long_url[0],
		ExpiresAt:   nowPlusMonth,
		CreatedAt:   now,
	}
	fmt.Println("urlData", urlData)

	preparedData, err := attributevalue.MarshalMap(urlData)
	fmt.Println("preparedData", preparedData)

	if err != nil {
		log.Fatalf("Failed to prepare data for database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to prepare data for database: " + err.Error(),
		})
	}

	_, err = basics.DynamoDbClient.PutItem(c.Context(), &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName),
		Item:      preparedData,
	})
	if err != nil {
		log.Fatalf("Failed to add item to ShortUrls table: %v\nURL Data: %+v\nPrepared Data: %+v", err, urlData, preparedData)

		// Handle specific error types if needed
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case "ProvisionedThroughputExceededException":
				log.Fatalf("Provisioned throughput exceeded for table: %s", basics.TableName)
				// Implement backoff or rate limiting strategy
			case "AccessDeniedException":
				log.Fatalf("Access denied for table: %s. Check IAM permissions.", basics.TableName)
			default:
				log.Fatalf("Unknown error: %v", awsErr)
			}
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add item to ShortUrls table: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"shortenURL": urlData.ShortUrl,
	})
}

func getURL(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"getURL": "hello",
	})
}

// // GetKey returns the composite primary key of the movie in a format that can be
// // sent to DynamoDB.
// func (movie Movie) GetKey() map[string]types.AttributeValue {
// 	title, err := attributevalue.Marshal(movie.Title)
// 	if err != nil {
// 		panic(err)
// 	}
// 	year, err := attributevalue.Marshal(movie.Year)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return map[string]types.AttributeValue{"title": title, "year": year}
// }

// // String returns the title, year, rating, and plot of a movie, formatted for the example.
// func (movie Movie) String() string {
// 	return fmt.Sprintf("%v\n\tReleased: %v\n\tRating: %v\n\tPlot: %v\n",
// 		movie.Title, movie.Year, movie.Info["rating"], movie.Info["plot"])
// }
