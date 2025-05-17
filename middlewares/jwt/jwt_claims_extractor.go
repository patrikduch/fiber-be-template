package jwt

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"fiber-be-template/utils/authctx"
)

func ExtractClaimsToContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenData := c.Locals("user")

		log.Println("[JWT] Raw token data from Locals:", tokenData)

		token, ok := tokenData.(*jwt.Token)
		if !ok {
			log.Println("[JWT] Token is not of type *jwt.Token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid JWT token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("[JWT] Token claims are not of type jwt.MapClaims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid JWT claims",
			})
		}

		log.Println("[JWT] Extracted claims:", claims)

		emailStr, emailOk := claims["email"].(string)
		if !emailOk || emailStr == "" {
			log.Println("[JWT] Missing or invalid 'email' claim")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing 'email' claim in token",
			})
		}

		log.Println("[JWT] email:", emailStr)

		// Store email in context
		ctx := authctx.WithUserEmail(c.UserContext(), emailStr)
		c.SetUserContext(ctx)

		log.Println("[JWT] Injected email into request context")

		return c.Next()
	}
}
