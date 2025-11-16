package routes

import (
	"axon/server/internal/config"

	"github.com/gofiber/fiber/v3"
)

// @ref https://restfulapi.net/resource-naming/
func Register(app *fiber.App, cfg *config.Config) {
	// TODO: include middleware api := app.Group("/api", middleware)
	api := app.Group("/api")
	v1 := api.Group("/v1")

	RegisterHealthV1(v1)
}
