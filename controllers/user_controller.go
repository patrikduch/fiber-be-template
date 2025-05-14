package controllers

import (


    "github.com/gofiber/fiber/v2"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/services/users"

)

// GetUsers godoc
// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {array} responses.UserResponseDto
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
    result, err := users.GetAllUsers()
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
// @Success 200 {object} responses.UserResponseDto "User details"
// @Failure 400 {object} map[string]string "Invalid UUID format"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
    id := c.Params("id") // this is a string
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

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user from request body
// @Tags users
// @Accept json
// @Produce json
// @Param user body requests.CreateUserRequestDto true "User to create"
// @Success 201 {object} responses.UserResponseDto
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
    var req requests.CreateUserRequestDto
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    user, err := users.CreateUser(req)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(user)
}
