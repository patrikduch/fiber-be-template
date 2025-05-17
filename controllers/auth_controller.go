package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"

	"fiber-be-template/commands/users/login_user"
	"fiber-be-template/dtos/users/requests"
	"fiber-be-template/queries/get_authenticated_user"
	"fiber-be-template/dtos/common"
)

var loginUserHandler = login_user.NewHandler()
var getAuthenticatedUserHandler = get_authenticated_user.NewHandler()

// LoginUser godoc
// @Summary Authenticate user
// @Description Login user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body requests.LoginUserRequestDto true "User credentials"
// @Success 200 {object} responses.LoginUserResponseDto
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
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

// Me godoc
// @Summary Get current user
// @Description Get authenticated user's information
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.AuthMeResponseDto
// @Failure 401 {object} common.Error401Response "Unauthorized"
// @Failure 500 {object} common.Error500Response "Internal server error"
// @Router /api/auth/me [get]
func Me(c *fiber.Ctx) error {
	user, err := getAuthenticatedUserHandler.Handle(c.UserContext(), get_authenticated_user.Query{})
	if err != nil {
		if err.Error() == "unauthorized" {
			return c.Status(401).JSON(common.Error401Response{
				Error: "Unauthorized",
			})
		}
		return c.Status(500).JSON(common.Error500Response{
			Error: err.Error(),
		})
	}
	return c.JSON(user)
}