package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-redis/redis/v8"
	"context"
	"go.uber.org/zap"
)

var db *dynamodb.DynamoDB
var logger *zap.Logger

func init() {
	// Our logger
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}

	// Our DynamoDB session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		log.Fatalf("Can't initialize db session: %v", err)
	}
	db = dynamodb.New(sess)
}

func main(){
	app := fiber.New()

	// Routes
	app.Post("/shorten", shortenURL)
	app.Get("/:shortCode", getURL)

	// Start server
	log.Fatal(app.Listen(":3000"))
}