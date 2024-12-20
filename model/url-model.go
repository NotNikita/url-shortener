package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// For DynamoDB level
type DBUrlData struct {
	ShortUrl    string `dynamodbav:"shortCode"`
	OriginalUrl string `dynamodbav:"originalURL"`
	CreatedAt   string `dynamodbav:"createdAt"`
	ExpiresAt   string `dynamodbav:"expiresAt"`
}

// For Frontend and Service level
type ViewUrlData struct {
	ShortUrl    string `json:"shortUrl" validate:"required"`
	OriginalUrl string `json:"originUrl" validate:"required"`
	CreatedAt   string `json:"createdAt" validate:"required"`
	ExpiresAt   string `json:"expiresAt" validate:"required"`
}

func (urlData DBUrlData) GeyKey() map[string]types.AttributeValue {
	shortUrl, err := attributevalue.Marshal(urlData.ShortUrl)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{
		"shortCode": shortUrl,
	}
}
