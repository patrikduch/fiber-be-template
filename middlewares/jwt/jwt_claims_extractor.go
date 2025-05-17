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

		subStr, ok := claims["sub"].(string)
		if !ok {
			log.Println("[JWT] Missing or invalid 'sub' claim")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing 'sub' claim in token",
			})
		}
		log.Println("[JWT] sub (user ID):", subStr)

		emailStr, _ := claims["email"].(string)
		log.Println("[JWT] email:", emailStr)

		ctx := authctx.WithUserID(c.UserContext(), subStr)
		if emailStr != "" {
			ctx = authctx.WithUserEmail(ctx, emailStr)
		}

		c.SetUserContext(ctx)
		log.Println("[JWT] Injected sub and email into request context")

		return c.Next()
	}
}
