package controllers

import (
    "context"

    "github.com/gofiber/fiber/v2"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/queries/get_all_users"
    "fiber-be-template/queries/get_user_by_email"
    "fiber-be-template/services/users"
    "fiber-be-template/commands/users/register_user"
)

var getAllUsersHandler = get_all_users.NewHandler()
var getUserByEmailHandler = get_user_by_email.NewHandler()

// GetUsers godoc
// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {array} responses.UserResponseDto
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
    result, err := getAllUsersHandler.Handle(context.Background(), get_all_users.Query{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(result)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Returns a single user based on their UUID
// @Tags users
// @Produce json
// @Param id path string true "User ID (UUID format)" format(uuid)
// @Success 200 {object} responses.UserResponseDto
// @Failure 400 {object} map[string]string "Invalid UUID format"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
    id := c.Params("id")
    user, err := users.GetUserByID(id)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    if user == nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }
    return c.JSON(user)
}

// GetUserByEmail godoc
// @Summary Get a user by email
// @Description Returns a single user based on their email
// @Tags users
// @Produce json
// @Param email query string true "User email"
// @Success 200 {object} responses.UserResponseDto
// @Failure 400 {object} map[string]string "Missing or invalid email"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/users/by-email [get]
func GetUserByEmail(c *fiber.Ctx) error {
    email := c.Query("email")
    if email == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Missing required email query parameter",
        })
    }

    user, err := getUserByEmailHandler.Handle(context.Background(), get_user_by_email.Query{Email: email})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    if user == nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    return c.JSON(user)
}


// RegisterUser godoc
// @Summary Register a new user
// @Description Creates a user with name, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body requests.RegisterUserRequestDto true "User registration payload"
// @Success 201 {object} responses.UserResponseDto
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users/register [post]
func RegisterUser(c *fiber.Ctx) error {
    var registerUserHandler = register_user.NewHandler()
    var req requests.RegisterUserRequestDto
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    user, err := registerUserHandler.Handle(context.Background(), register_user.Command{Payload: req})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(user)
}