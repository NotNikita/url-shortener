package service

import (
	"context"
	"time"

	utility_service "url-shortener/service/utility"
	"url-shortener/store"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

type ServiceManager struct {
	UrlsService  *UrlsWebService
	HashService  *utility_service.HashingService
	CacheService *utility_service.CachingService
}

// NewServiceManager creates new service manager
func NewServiceManager(ctx context.Context, store *store.Store) (*ServiceManager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	// Init memcache
	// Expiration: 5 minutes, Purges expired items every 10 minutes
	urlMemcache := cache.New(5*time.Minute, 10*time.Minute)

	hashService := utility_service.NewHashingService(ctx)
	cacheService := utility_service.NewCachingService(ctx, urlMemcache)

	return &ServiceManager{
		UrlsService:  NewUrlsWebService(ctx, store, hashService, cacheService),
		HashService:  hashService,
		CacheService: cacheService,
	}, nil
}
