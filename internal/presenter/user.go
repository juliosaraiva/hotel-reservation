package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
)

func UserResponseSuccess(data []*models.User) *fiber.Map {
	return &fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
		"error":  "",
	}
}

func UserResponseError(err error) *fiber.Map {
	return &fiber.Map{
		"status": fiber.StatusBadRequest,
		"data":   "",
		"error":  err.Error(),
	}
}
