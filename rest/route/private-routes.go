package route

import (
	"url-shortener/rest/controller"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PrivateRoutes(a *fiber.App, urls *controller.UrlsController) {
	group := a.Group("/api/v1")

	// Url related routes:
	group.Post("/shorten", urls.ShortenUrl)
	group.Get("/:shortCode", urls.GetOriginalUrl)
}
