package main

import (
	"context"
	"fmt"
	"log"

	"url-shortener/config"
	"url-shortener/logger"
	"url-shortener/rest/controller"
	"url-shortener/rest/middleware"
	"url-shortener/rest/route"
	"url-shortener/service"
	"url-shortener/store"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	// "github.com/go-redis/redis/v8"
)

// TableBasics encapsulates the Amazon DynamoDB service actions.
// Contains a DynamoDB service client that is used to act on the specified table.
type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

var basics TableBasics

const (
	HASH_LENGTH = 6
)

func main() {
	ctx := context.Background()
	// config + logger
	config.Get()
	fiberConf := config.GetFiberConfig()
	l := logger.Get()

	// Init repository store
	store, err := store.NewStore(ctx)
	if err != nil {
		l.Fatal("Failed to initialize store", zap.Error(err))
	}

	// Init service manager
	serviceManager, err := service.NewServiceManager(ctx, store)
	if err != nil {
		l.Fatal("Failed to initialize service manager", zap.Error(err))
	}

	// Init controllers
	urlsController := controller.NewUrlsController(
		ctx, serviceManager, l,
	)

	// Init fiber instance
	app := fiber.New(fiberConf)

	// Middlewares.
	middleware.FiberMiddleware(app)

	fmt.Println("Server started")

	// Routes
	route.SwaggerRoutes(app)
	route.PublicRoutes(app)
	route.PrivateRoutes(app, urlsController)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
