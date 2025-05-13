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

func GetUsers(c *fiber.Ctx) error {
    return c.JSON(users)
}

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

func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
    }

    user.ID = len(users) + 1
    users = append(users, user)

    return c.Status(201).JSON(user)
}
