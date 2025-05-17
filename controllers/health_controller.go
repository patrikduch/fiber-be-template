package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"fiber-be-template/database"
	"fiber-be-template/dtos/common"
	"fiber-be-template/dtos/health/responses"
)

// GetHealth godoc
// @Summary Health check
// @Description Returns application and database health status
// @Tags Health
// @Produce json
// @Success 200 {object} responses.HealthStatusResponse
// @Failure 500 {object} common.Error500Response "Internal server error"
// @Router /api/health [get]
func GetHealth(c *fiber.Ctx) error {
	if err := database.DB.Ping(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(common.Error500Response{
			Error: "Database unreachable",
		})
	}

	return c.JSON(responses.HealthStatusResponse{
		App: "ok",
		DB:  "connected",
	})
}
