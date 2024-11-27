package aws

import (
	"url-shortener/model"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
)

type UrlsRepo struct {
	awsClient *dynamodb.Client
}

// create UrlsRepo
func NewUrlsRepo(client *dynamodb.Client) *UrlsRepo {
	return &UrlsRepo{
		awsClient: client,
	}
}

func (repo *UrlsRepo) GetUrl(ctx *fiber.Ctx, shortUrl string) *model.UrlData {
	// TODO: move logic from controller
	panic("Not implemented")
}

func (repo *UrlsRepo) CreateUrl(ctx *fiber.Ctx, obj *model.UrlData) (*model.UrlData, error) {
	// TODO: move logic from controller
	panic("Not implemented")
}

// TODO:
func (repo *UrlsRepo) UpdateUrl(ctx *fiber.Ctx, obj *model.UrlData) (*model.UrlData, error) {
	panic("Not implemented")
}

func (repo *UrlsRepo) DeleteUrl(ctx *fiber.Ctx, shortUrl string) error {
	panic("Not implemented")
}
