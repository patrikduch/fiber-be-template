package controllers

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "fiber-be-template/database"
)

// GetHealth godoc
// @Summary Health check
// @Description Returns app and DB status
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/health [get]
func GetHealth(c *fiber.Ctx) error {
    if err := database.DB.Ping(); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "status": "unhealthy",
            "db":     "unreachable",
        })
    }

    return c.JSON(fiber.Map{
        "status": "ok",
        "db":     "connected",
    })
}
