package routes

import (
    "github.com/gofiber/fiber/v2"
    "fiber-be-template/controllers"
)

func RegisterHealthRoutes(app fiber.Router) {
    app.Get("/api/health", controllers.GetHealth)
}
