package routes

import (
    "github.com/gofiber/fiber/v2"
    "fiber-be-template/controllers"
)

func RegisterUserRoutes(app fiber.Router) {
    user := app.Group("/api/users")
    user.Get("/", controllers.GetUsers)
    user.Get("/:id", controllers.GetUserByID)
    user.Post("/", controllers.CreateUser)
}
