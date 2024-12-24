package route

import (
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	group := a.Group("/api/v1")

	// Routes for GET method:
	group.Get("/ping", healthRoute)
}

// @Summary Checks if server is running
// @Description Should return body with status of "pong"
// @Tags Public
// @Produce json
// @Success 200
// @Router /ping [get]
func healthRoute(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "pong",
	})
}
