package controllers

import (
    "strconv"

    "github.com/gofiber/fiber/v2"

    "fiber-be-template/models"
    "fiber-be-template/dtos/users/requests"
    "fiber-be-template/dtos/users/responses"
    userMappers "fiber-be-template/mappers/users"
)

var usersStore = []models.User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// GetUsers godoc
// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {array} responses.UserResponseDto
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
    var result []responses.UserResponseDto
    for _, u := range usersStore {
        result = append(result, userMappers.ToUserResponseDto(u))
    }
    return c.JSON(result)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Returns a single user based on ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} responses.UserResponseDto
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    for _, u := range usersStore {
        if u.ID == id {
            return c.JSON(userMappers.ToUserResponseDto(u))
        }
    }

    return c.Status(404).JSON(fiber.Map{"error": "User not found"})
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
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
    var req requests.CreateUserRequestDto
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    newUser := userMappers.ToUserModel(req, len(usersStore)+1)
    usersStore = append(usersStore, newUser)

    return c.Status(201).JSON(userMappers.ToUserResponseDto(newUser))
}
