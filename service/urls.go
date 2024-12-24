package service

import (
	"context"

	"url-shortener/model"
	"url-shortener/store"

	"github.com/pkg/errors"
)

type UrlsWebService struct {
	ctx            context.Context
	store          *store.Store
	hashingService *HashingService
}

func NewUrlsWebService(ctx context.Context, store *store.Store, hashingService *HashingService) *UrlsWebService {
	return &UrlsWebService{
		ctx,
		store,
		hashingService,
	}
}

// Retrieve original URL
func (service *UrlsWebService) GetUrl(ctx context.Context, shortUrl string) (*model.ViewUrlData, error) {
	dbUrlData, err := service.store.Urls.GetUrl(ctx, shortUrl)
	if err != nil {
		return nil, errors.Wrap(err, "service.urls.GetUrl")
	}

	return dbUrlData.ToView(), nil
}

// Shorten original URL
func (service *UrlsWebService) CreateUrl(ctx context.Context, obj *model.ViewUrlData) (*model.ViewUrlData, error) {
	shortCode, err := service.hashingService.GenerateXXHash3BasedOnOriginURL(ctx, obj.OriginalUrl)
	if err != nil {
		return nil, errors.Wrap(err, "service.urls.CreateUrl")
	}

	obj.ShortUrl = shortCode
	dbUrlData, err := service.store.Urls.CreateUrl(ctx, obj.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "service.urls.CreateUrl")
	}

	return dbUrlData.ToView(), nil
}

// TODO:
// Update existing short URL to point to another Origin
func (service *UrlsWebService) UpdateUrl(ctx context.Context, obj *model.ViewUrlData) (*model.ViewUrlData, error) {
	panic("Not implemented")
}

// Delete existing short URL
func (service *UrlsWebService) DeleteUrl(ctx context.Context, shortUrl string) error {
	panic("Not implemented")
}
