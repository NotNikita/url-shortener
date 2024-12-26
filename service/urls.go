package service

import (
	"context"

	"url-shortener/model"
	utility_service "url-shortener/service/utility"
	"url-shortener/store"

	"github.com/pkg/errors"
)

type UrlsWebService struct {
	ctx            context.Context
	store          *store.Store
	hashingService *utility_service.HashingService
	cachingService *utility_service.CachingService
}

func NewUrlsWebService(ctx context.Context, store *store.Store, hashingService *utility_service.HashingService, cacheService *utility_service.CachingService) *UrlsWebService {
	return &UrlsWebService{
		ctx,
		store,
		hashingService,
		cacheService,
	}
}

// Retrieve original URL
func (service *UrlsWebService) GetUrl(ctx context.Context, shortUrl string) (*model.ViewUrlData, error) {
	// Checking cache before accessing DB:
	if cacheRes, ok := service.cachingService.GetCache(shortUrl); ok {
		return &model.ViewUrlData{
			OriginalUrl: cacheRes,
		}, nil
	}

	// Cache miss, accessing DB:
	dbUrlData, err := service.store.Urls.GetUrl(ctx, shortUrl)
	if err != nil {
		return nil, errors.Wrap(err, "service.urls.GetUrl")
	}

	service.cachingService.UpdateCache(shortUrl, dbUrlData.OriginalURL)
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

	service.cachingService.UpdateCache(dbUrlData.ShortCode, dbUrlData.OriginalURL)
	return dbUrlData.ToView(), nil
}

// Update existing short URL to point to another Origin
func (service *UrlsWebService) UpdateUrl(ctx context.Context, obj *model.ViewUrlData) (*model.ViewUrlData, error) {
	dbUrlData, err := service.store.Urls.UpdateUrl(ctx, obj.ToDB())
	if err != nil {
		return nil, errors.Wrap(err, "service.urls.UpdateUrl")
	}

	service.cachingService.UpdateCache(dbUrlData.ShortCode, dbUrlData.OriginalURL)
	return dbUrlData.ToView(), nil
}

// Delete existing short URL
func (service *UrlsWebService) DeleteUrl(ctx context.Context, shortUrl string) error {
	err := service.store.Urls.DeleteUrl(ctx, shortUrl)
	if err != nil {
		return errors.Wrap(err, "service.urls.DeleteUrl")
	}

	service.cachingService.EjectCache(shortUrl)
	return nil
}
