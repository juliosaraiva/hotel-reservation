package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
)

func HotelSuccessResponse(data []*models.HotelRooms) *fiber.Map {
	return &fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
		"error":  "",
	}
}

func HotelErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func RoomSuccessResponse(data []*models.Room) *fiber.Map {
	return &fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
		"error":  "",
	}
}

func RoomErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
