package utility

import (
	"context"

	"github.com/patrickmn/go-cache"
)

type CachingService struct {
	ctx      context.Context
	memcache *cache.Cache
}

func NewCachingService(ctx context.Context, memcache *cache.Cache) *CachingService {
	return &CachingService{
		ctx:      ctx,
		memcache: memcache,
	}
}

func (c *CachingService) GetCache(shortCode string) (string, bool) {
	if cacheVal, ok := c.memcache.Get(shortCode); ok {
		return cacheVal.(string), true
	}

	return "", false
}

func (c *CachingService) UpdateCache(shortUrl string, originUrl string) {
	c.memcache.Set(shortUrl, originUrl, cache.DefaultExpiration)
}

func (c *CachingService) EjectCache(shortUrl string) {
	c.memcache.Delete(shortUrl)
}
