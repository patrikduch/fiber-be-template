package jwt

import (
    "os"

    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v3"
)

func Protected() fiber.Handler {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        panic("JWT_SECRET not set")
    }

    return jwtware.New(jwtware.Config{
        SigningKey:   []byte(secret),
        ContextKey:   "user", // Token payload available via c.Locals("user")
        ErrorHandler: errorHandler,
    })
}

func errorHandler(c *fiber.Ctx, err error) error {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Unauthorized",
    })
}
