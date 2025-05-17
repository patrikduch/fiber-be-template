package roles

import (
	"github.com/gofiber/fiber/v2"
	"fiber-be-template/utils/authctx"
	"fiber-be-template/constants"
)

func RequireAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := authctx.GetAuthenticatedUserWithRole(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		if user.Role == nil || user.Role.Name != constants.RoleAdmin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "You do not have permission to perform this action.",
			})
		}

		return c.Next()
	}
}
