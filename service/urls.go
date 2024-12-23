package service

import (
	"context"
	"url-shortener/model"
	"url-shortener/store"
)

type UrlsWebService struct {
	ctx   context.Context
	store *store.Store
}

func NewUrlsWebService(ctx context.Context, store *store.Store) *UrlsWebService {
	return &UrlsWebService{
		ctx,
		store,
	}
}

// GetUrl
func (service *UrlsWebService) GetUrl(ctx context.Context, shortUrl string) *model.ViewUrlData {
	// TODO: move logic from controller
	panic("Not implemented")
}

func (service *UrlsWebService) CreateUrl(ctx context.Context, obj *model.ViewUrlData) (*model.ViewUrlData, error) {
	// TODO: move logic from controller
	panic("Not implemented")
}

// TODO:
func (service *UrlsWebService) UpdateUrl(ctx context.Context, obj *model.ViewUrlData) (*model.ViewUrlData, error) {
	panic("Not implemented")
}

func (service *UrlsWebService) DeleteUrl(ctx context.Context, shortUrl string) error {
	panic("Not implemented")
}
