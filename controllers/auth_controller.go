package controllers

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "fiber-be-template/commands/users/login_user"
    "fiber-be-template/dtos/users/requests"
)

var loginUserHandler = login_user.NewHandler()

// LoginUser godoc
// @Summary Authenticate user
// @Description Login user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body requests.LoginUserRequestDto true "User credentials"
// @Success 200 {object} responses.LoginUserResponseDto
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/login [post]
func LoginUser(c *fiber.Ctx) error {
    var req requests.LoginUserRequestDto
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    user, err := loginUserHandler.Handle(context.Background(), login_user.Command{Payload: req})
    if err != nil {
        if err.Error() == "invalid credentials" {
            return c.Status(401).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(user)
}
