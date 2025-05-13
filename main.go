package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/cors"

    "fiber-be-template/routes"
)

func main() {
    app := fiber.New()

    // Middleware
    app.Use(logger.New())
    app.Use(cors.New())

    // Health route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber!")
    })

    app.Get("/api/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })

    // Register user routes
    routes.RegisterUserRoutes(app)

    // Start server
    log.Fatal(app.Listen(":3000"))
}
