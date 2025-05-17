package roles

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"fiber-be-template/utils/authctx"
)

func RequireRoles(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := authctx.GetAuthenticatedUserWithRole(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: user could not be resolved",
			})
		}

		if user.Role == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": fmt.Sprintf(
					"You do not have permission to perform this action. This action requires one of the following roles: [%s].",
					strings.Join(allowedRoles, ", "),
				),
			})
		}

		for _, allowed := range allowedRoles {
			if user.Role.Name == allowed {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf(
				"You do not have permission to perform this action. Your role '%s' is not authorized.",
				user.Role.Name,
			),
		})
	}
}
