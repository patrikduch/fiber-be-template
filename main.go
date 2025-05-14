// @title Fiber BE Template API
// @version 1.0
// @description Fiber backend template
// @BasePath /
package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/cors"

    "fiber-be-template/routes"

	"github.com/gofiber/swagger" // 👈 Swagger UI handler
    _ "fiber-be-template/docs" // 👈 Generated docs
	"github.com/joho/godotenv"
	"fiber-be-template/database"
)

func main() {
    app := fiber.New()

 	// Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  .env file not found, using system env vars")
    }

    // Initialize DB
    database.Init()

    // Middleware
    app.Use(logger.New())
    app.Use(cors.New())

	   // Swagger UI
    app.Get("/swagger/*", swagger.HandlerDefault)

    // Health route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber!")
    })

    // Register user routes
    routes.RegisterUserRoutes(app);
	routes.RegisterHealthRoutes(app);
    routes.RegisterAuthRoutes(app);

    // Start server
    log.Fatal(app.Listen(":3000"))
}
