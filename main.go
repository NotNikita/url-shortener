package main

import (
	"fmt"
	"log"

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
