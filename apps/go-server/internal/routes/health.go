package routes

import (
	"axon/server/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterHealthV1(router fiber.Router) {
	r := router.Group("/health")

	r.Get("/", handlers.GetHealth)
}
