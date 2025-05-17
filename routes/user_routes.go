package routes

import (
	"github.com/gofiber/fiber/v2"
	"fiber-be-template/controllers"
	jwtmiddleware "fiber-be-template/middlewares/jwt"
	rolemiddleware "fiber-be-template/middlewares/roles"
)

func RegisterUserRoutes(app fiber.Router) {
	user := app.Group("/api/users")

	// Public route
	user.Post("/register", controllers.RegisterUser)

	// Apply base JWT middleware (token validation + claim injection)
	user.Use(
		jwtmiddleware.Protected(),
		jwtmiddleware.ExtractClaimsToContext(),
	)

	// 📧 Get user by email — admin only
	user.Get("/by-email",
		rolemiddleware.RequireRoles("admin"),
		controllers.GetUserByEmail,
	)

	// 📋 Get all users — admin only
	user.Get("/",
		rolemiddleware.RequireRoles("admin"),
		controllers.GetUsers,
	)

	// 🔍 Get user by ID — admin only
	user.Get("/:id",
		rolemiddleware.RequireRoles("admin"),
		controllers.GetUserByID,
	)
}
