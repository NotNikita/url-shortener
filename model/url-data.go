package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UrlData struct {
	ShortUrl    string `dynamodbav:"shortCode"`
	OriginalUrl string `dynamodbav:"originalURL"`
	CreatedAt   string `dynamodbav:"createdAt"`
	ExpiresAt   string `dynamodbav:"expiresAt"`
}

func (urlData UrlData) GeyKey() map[string]types.AttributeValue {
	shortUrl, err := attributevalue.Marshal(urlData.ShortUrl)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{
		"shortCode": shortUrl,
	}
}
