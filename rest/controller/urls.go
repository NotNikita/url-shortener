package controller

import (
	"context"
	"fmt"
	"log"
	"time"
	"url-shortener/logger"
	"url-shortener/model"
	"url-shortener/service"
	"url-shortener/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gofiber/fiber/v2"
)

// Urls Controller
type UrlsController struct {
	ctx      context.Context
	services *service.ServiceManager
	logger   *logger.Logger
}

// Creates a new Urls Controller
func NewUrlsController(ctx context.Context, services *service.ServiceManager, logger *logger.Logger) *UrlsController {
	return &UrlsController{
		ctx:      ctx,
		services: services,
		logger:   logger,
	}
}

// Interface of UrlsController: shorten a given url
func (cnt *UrlsController) ShortenUrl(c *fiber.Ctx) error {
	var newUrlData model.ViewUrlData
	form, err := c.MultipartForm()
	if err != nil {
		log.Fatalf("Error parsing provided data for shortenURL: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form: " + err.Error(),
		})
	}

	long_url := form.Value["long_url"]
	if len(long_url) == 0 || !utils.IsValidUrl(long_url[0]) {
		log.Fatalf("Error when validating provided url to shorten")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to validate url: " + err.Error(),
		})
	}

	now := time.Now().Format(time.RFC3339)
	nowPlusMonth := time.Now().AddDate(0, 1, 0).Format(time.RFC3339)
	// put valid in db
	// TODO: short data xxHash creation
	urlData := model.ViewUrlData{
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

// Interface of UrlsController: get original url for provided short and redirect
func (cnt *UrlsController) GetOriginalUrl(c *fiber.Ctx) error {
	param := c.Params("shortCode")
	if len(param) < 1 { // TODO: HASH_LENGTH {
		log.Printf("Error when retrieving shortCode from request url")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve shortCode in getURL function",
		})
	}

	urlData := model.UrlData{
		ShortUrl: param,
	}
	response, err := basics.DynamoDbClient.GetItem(c.Context(), &dynamodb.GetItemInput{
		TableName: aws.String(basics.TableName),
		Key:       urlData.GeyKey(),
	})
	if err != nil {
		log.Printf("Error when retrieving urlData from db in getURL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to retrieve urlData from DB in getURL function",
			})
	}

	if response.Item == nil {
		log.Printf("No item found in DB for shortCode: %s", param)
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"error": "No URL found for the provided shortCode",
			})
	}

	err = attributevalue.UnmarshalMap(response.Item, &urlData)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
	}

	return c.Redirect(urlData.OriginalUrl, fiber.StatusFound)
}
