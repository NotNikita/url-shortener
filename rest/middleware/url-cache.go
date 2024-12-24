package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

// Cache for retrieving origin URL by provided shortCode
func VerifyGetOriginalUrl(memcache *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortCode := c.Params("shortCode")
		if cacheVal, ok := memcache.Get(shortCode); ok {
			fmt.Println("Retrieving cache", cacheVal.(string))
			return c.Redirect(cacheVal.(string), fiber.StatusFound)
		}

		return c.Next()
	}
}

func PutOriginalUrlInCache(memcache *cache.Cache, shortUrl string, originUrl string) {
	memcache.Set(shortUrl, originUrl, cache.DefaultExpiration)
}