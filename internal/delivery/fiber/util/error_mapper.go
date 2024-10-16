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
	case "not.found.error":
		return fiber.StatusNotFound
	default:
		return fiber.StatusInternalServerError
	}
}
