package helper

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func JsonResponse(c *fiber.Ctx, data interface{}) {
	if err := c.Status(http.StatusOK).JSON(fiber.Map{
		"data": data,
	}); err != nil {
		err := c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		if err != nil {
			return
		}
	}
}
