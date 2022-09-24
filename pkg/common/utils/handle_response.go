package utils

import (
	"delta-go/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func HandleResponse(c *fiber.Ctx, statusCode int, message string) error {
	response := models.Response{
		StatusCode: statusCode,
		Message:    message,
	}
	return c.Status(response.StatusCode).JSON(&response)
}
