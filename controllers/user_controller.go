package controllers

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/queries/get_all_users"
    "fiber-be-template/queries/get_user_by_email"
    "fiber-be-template/queries/get_user_by_id"
    "fiber-be-template/commands/users/register_user"
    "fiber-be-template/dtos/common"
)

var (
	getAllUsersHandler     = get_all_users.NewHandler()
	getUserByEmailHandler  = get_user_by_email.NewHandler()
	getUserByIDHandler     = get_user_by_id.NewHandler()
)

// GetUsers godoc
// @Summary      Get all users
// @Description  Retrieves a list of all users.
// @Tags         Users
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} responses.UserResponseDto
// @Failure      401 {object} common.Error401Response "Unauthorized"
// @Failure      500 {object} common.Error500Response "Internal Server Error"
// @Router       /api/users [get]
func GetUsers(c *fiber.Ctx) error {
    result, err := getAllUsersHandler.Handle(context.Background(), get_all_users.Query{})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(common.Error500Response{
            Error: "Failed to retrieve users",
        })
    }

    return c.Status(fiber.StatusOK).JSON(result)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Returns a single user based on their UUID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID (UUID format)" format(uuid)
// @Success 200 {object} responses.UserResponseDto
// @Failure 400 {object} common.Error400Response "Invalid UUID format"
// @Failure 401 {object} common.Error401Response "Unauthorized"
// @Failure 404 {object} common.Error404Response "User not found"
// @Failure 500 {object} common.Error500Response "Internal server error"
// @Router /api/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Optional: validate UUID format early
	if len(id) != 36 {
		return c.Status(fiber.StatusBadRequest).JSON(common.Error400Response{
			Error: "Invalid UUID format",
		})
	}

	user, err := getUserByIDHandler.Handle(context.Background(), get_user_by_id.Query{ID: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.Error500Response{
			Error: "Failed to retrieve user: " + err.Error(),
		})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(common.Error404Response{
			Error: "User not found",
		})
	}

	return c.JSON(user)
}

// GetUserByEmail godoc
// @Summary Get a user by email
// @Description Returns a single user based on their email
// @Tags Users
// @Produce json
// @Param email query string true "User email"
// @Security BearerAuth
// @Success 200 {object} responses.UserResponseDto
// @Failure 400 {object} common.Error400Response "Bad request"
// @Failure 401 {object} common.Error401Response "Unauthorized"
// @Failure 404 {object} common.Error404Response "User not found"
// @Failure 500 {object} common.Error500Response "Internal server error"
// @Router /api/users/by-email [get]
func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.Error400Response{
			Error: "Missing required 'email' query parameter",
		})
	}

	user, err := getUserByEmailHandler.Handle(context.Background(), get_user_by_email.Query{Email: email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.Error500Response{
			Error: "Failed to retrieve user",
		})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(common.Error404Response{
			Error: "User not found",
		})
	}

	return c.JSON(user)
}
// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a user with name, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body requests.RegisterUserRequestDto true "User registration payload"
// @Success 201 {object} responses.UserResponseDto
// @Failure 400 {object} common.Error400Response "Invalid input"
// @Failure 500 {object} common.Error500Response "Internal server error"
// @Router /api/users/register [post]
func RegisterUser(c *fiber.Ctx) error {
	var registerUserHandler = register_user.NewHandler()

	var req requests.RegisterUserRequestDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.Error400Response{
			Error: "Invalid request body",
		})
	}

	user, err := registerUserHandler.Handle(context.Background(), register_user.Command{Payload: req})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(common.Error500Response{
			Error: "Failed to register user: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}