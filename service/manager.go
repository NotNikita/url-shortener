package service

import (
	"context"

	"github.com/pkg/errors"

	"url-shortener/store"
)

type ServiceManager struct {
	UrlsService *UrlsWebService
	HashService *HashingService
}

// NewServiceManager creates new service manager
func NewServiceManager(ctx context.Context, store *store.Store) (*ServiceManager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}

	return &ServiceManager{
		UrlsService: NewUrlsWebService(ctx, store),
		HashService: NewHashingService(ctx),
	}, nil
}
