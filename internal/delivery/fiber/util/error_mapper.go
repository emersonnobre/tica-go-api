package util

import (
	"github.com/gofiber/fiber/v2"
)

func CoreErrorToHttpError(coreErrorName string) int {
	switch coreErrorName {
	case "internal.server.error":
		return fiber.StatusInternalServerError
	case "validation.error":
		return fiber.StatusBadRequest
	default:
		return fiber.StatusInternalServerError
	}
}
