package routes

import (
    "github.com/gofiber/fiber/v2"
    "fiber-be-template/controllers"
    jwtmiddleware "fiber-be-template/middlewares/jwt"
)

func RegisterUserRoutes(app fiber.Router) {
    user := app.Group("/api/users")

    // 🔐 JWT-protected routes
    user.Use(jwtmiddleware.Protected())

    // 📥 User registration (must come before dynamic :id route)
    user.Post("/register", controllers.RegisterUser)

    // 📧 Get user by email
    user.Get("/by-email", controllers.GetUserByEmail)

    // 📋 Get all users
    user.Get("/", controllers.GetUsers)

    // 🔍 Get by ID (must be last to avoid route conflicts)
    user.Get("/:id", controllers.GetUserByID)
}
