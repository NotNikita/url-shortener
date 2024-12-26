package aws

import (
	"context"
	"log"

	"url-shortener/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

func (repo *UrlsRepo) UpdateUrl(ctx context.Context, obj *model.DBUrlData) (*model.DBUrlData, error) {
	var err error
	var response *dynamodb.UpdateItemOutput
	var resultedUrlData model.DBUrlData

	update := expression.Set(expression.Name("originalURL"), expression.Value(obj.OriginalURL))
	update.Set(expression.Name("expiresAt"), expression.Value(obj.ExpiresAt))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
		return nil, errors.Wrap(err, "Couldn't build expression for update")
	}

	response, err = repo.awsClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(URLS_TABLE_NAME),
		Key:                       obj.ToView().GeyKey(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		log.Printf("Couldn't update origin url of %v. Here's why: %v\n", obj.ShortCode, err)
		return nil, errors.Wrap(err, "Couldn't update origin url of "+obj.ShortCode)
	}

	err = attributevalue.UnmarshalMap(response.Attributes, &resultedUrlData)
	if err != nil {
		log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		return nil, errors.Wrap(err, "Couldn't unmarshal response")
	}
	return &resultedUrlData, nil
}

func (repo *UrlsRepo) DeleteUrl(ctx context.Context, shortUrl string) error {
	_, err := repo.awsClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(URLS_TABLE_NAME),
		Key:       model.MarshalShortUrl(shortUrl),
	})
	if err != nil {
		log.Printf("Couldn't delete url of %v. Here's why: %v\n", shortUrl, err)
		return errors.Wrap(err, "Couldn't delete url of "+shortUrl)
	}
	return nil
}
