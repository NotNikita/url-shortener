package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// For DynamoDB level
type DBUrlData struct {
	ShortCode   string `dynamodbav:"shortCode"`
	OriginalURL string `dynamodbav:"originalURL"`
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

// Converts ViewUrlData to DBUrlData
func (urlData *ViewUrlData) ToDB() *DBUrlData {
	return &DBUrlData{
		ShortCode:   urlData.ShortUrl,
		OriginalURL: urlData.OriginalUrl,
		CreatedAt:   urlData.CreatedAt,
		ExpiresAt:   urlData.ExpiresAt,
	}
}

// Converts DBUrlData to ViewUrlData
func (urlDB *DBUrlData) ToView() *ViewUrlData {
	return &ViewUrlData{
		ShortUrl:    urlDB.ShortCode,
		OriginalUrl: urlDB.OriginalURL,
		CreatedAt:   urlDB.CreatedAt,
		ExpiresAt:   urlDB.ExpiresAt,
	}
}

func (urlData ViewUrlData) GeyKey() map[string]types.AttributeValue {
	return MarshalShortUrl(urlData.ShortUrl)
}

func MarshalShortUrl(url string) map[string]types.AttributeValue {
	shortUrl, err := attributevalue.Marshal(url)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{
		"shortCode": shortUrl,
	}
}
