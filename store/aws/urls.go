package aws

import (
	"context"
	"log"

	"url-shortener/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
)

const (
	URLS_TABLE_NAME = "ShortUrls"
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

func (repo *UrlsRepo) GetUrl(ctx context.Context, shortUrl string) (*model.DBUrlData, error) {
	var resultedUrlData model.DBUrlData
	urlDataKey := model.MarshalShortUrl(shortUrl)
	awsRepoObject, err := repo.awsClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(URLS_TABLE_NAME),
		Key:       urlDataKey,
	})
	if err != nil {
		log.Printf("Error when retrieving urlData from db in getURL: %v", err)
		return nil, errors.Wrap(err, "Failed to retrieve urlData from DB in getURL function")
	}

	if awsRepoObject.Item == nil {
		log.Printf("No item found in DB for shortCode: %s, %s", shortUrl, urlDataKey)
		return nil, errors.Wrap(err, "No URL found for the provided shortCode")
	}

	err = attributevalue.UnmarshalMap(awsRepoObject.Item, &resultedUrlData)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		return nil, errors.Wrap(err, "Couldn't unmarshal response")
	}
	return &resultedUrlData, nil
}

func (repo *UrlsRepo) CreateUrl(ctx context.Context, obj *model.DBUrlData) (*model.DBUrlData, error) {
	preparedData, err := attributevalue.MarshalMap(obj)
	if err != nil {
		log.Printf("Failed to marshal DBUrlData")
		return nil, errors.Wrap(err, "Failed to marshal DBUrlData")
	}

	_, err = repo.awsClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:    aws.String(URLS_TABLE_NAME),
		Item:         preparedData,
		ReturnValues: types.ReturnValueNone,
	})
	if err != nil {
		log.Printf("Failed to add item to ShortUrls table: %v\nURL Data: %+v\n", err, obj)
		return nil, errors.Wrap(err, "Failed to add item to ShortUrls table")
	}

	return obj, nil
}

// TODO:
func (repo *UrlsRepo) UpdateUrl(ctx context.Context, obj *model.DBUrlData) (*model.DBUrlData, error) {
	panic("Not implemented")
}

func (repo *UrlsRepo) DeleteUrl(ctx context.Context, shortUrl string) error {
	panic("Not implemented")
}
