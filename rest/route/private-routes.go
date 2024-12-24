package route

import (
	"url-shortener/rest/controller"
	"url-shortener/rest/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

// PublicRoutes func for describe group of public routes.
func PrivateRoutes(a *fiber.App, cache *cache.Cache, urls *controller.UrlsController) {
	group := a.Group("/api/v1")

	// Url related routes:
	group.Post("/shorten", urls.ShortenUrl)
	group.Get("/:shortCode", middleware.VerifyGetOriginalUrl(cache), urls.GetOriginalUrl)
	group.Put("/:shortCode", urls.UpdateUrl)
	group.Delete("/:shortCode", urls.DeleteUrl)
}
