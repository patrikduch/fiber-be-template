package roles

import (
	"github.com/gofiber/fiber/v2"
	"fiber-be-template/utils/authctx"
)

func RequireAuthenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := authctx.UserEmailFromContext(c.UserContext())
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "You must be logged in to access this resource.",
			})
		}
		return c.Next()
	}
}
