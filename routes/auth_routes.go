package routes

import (
    "github.com/gofiber/fiber/v2"
    "fiber-be-template/controllers"
)

func RegisterAuthRoutes(app fiber.Router) {
    auth := app.Group("/api/auth")

    auth.Post("/login", controllers.LoginUser)
}
