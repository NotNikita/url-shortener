package route

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger" // swagger handler
)

// https://github.com/gofiber/swagger?tab=readme-ov-file#canonical-example
func SwaggerRoutes(a *fiber.App) {
	group := a.Group("/swagger")

	// Routes for GET method:
	group.Get("*", swagger.HandlerDefault)
}
