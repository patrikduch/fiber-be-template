package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber-be-template/controllers"
	jwtmiddleware "fiber-be-template/middlewares/jwt"
)

func RegisterAuthRoutes(app fiber.Router) {
	auth := app.Group("/api/auth")

	// 🔓 Public route
	auth.Post("/login", controllers.LoginUser)

	// 🔐 Protected route
	auth.Use(jwtmiddleware.Protected())
    auth.Use(jwtmiddleware.ExtractClaimsToContext())  // ✅ To populate context

	auth.Get("/me", controllers.Me)
}
