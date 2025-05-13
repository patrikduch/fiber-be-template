package controllers

import (
    "strconv"
    "github.com/gofiber/fiber/v2"

	"fiber-be-template/models"
)

var users = []models.User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// GetUsers godoc
// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
    return c.JSON(users)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Returns a single user based on ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    for _, user := range users {
        if user.ID == id {
            return c.JSON(user)
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
// @Param user body models.User true "User to create"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /api/users [post]
func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    user.ID = len(users) + 1
    users = append(users, user)

    return c.Status(201).JSON(user)
}
