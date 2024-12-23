package route

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoutes(a *fiber.App) {
	group := a.Group("/swagger")

	// Routes for GET method:
	group.Get("*", swagger.HandlerDefault)
}
