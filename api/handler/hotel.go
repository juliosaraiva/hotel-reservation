package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
	"github.com/juliosaraiva/hotel-reservation/internal/presenter"
)

func GetHotels(hotel interfaces.Hotel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hotels, err := hotel.GetAll(c.Context())
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.HotelErrorResponse(err))
		}
		return c.JSON(presenter.HotelSuccessResponse(hotels))
	}
}

func GetHotelDetails(hotel interfaces.Hotel) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := hotel.Get(c.Context(), id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.HotelErrorResponse(err))
		}
		return c.JSON(presenter.HotelSuccessResponse(res))
	}
}

func GetRoom(room interfaces.Room) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := room.GetRoom(c.Context(), id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.RoomErrorResponse(err))
		}
		return c.JSON(presenter.RoomSuccessResponse(res))
	}
}

func GetRooms(room interfaces.Room) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := room.GetRooms(c.Context())
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.RoomErrorResponse(err))
		}
		return c.JSON(presenter.RoomSuccessResponse(res))
	}
}

func GetRoomsByHotelID(room interfaces.Room) fiber.Handler {
	return func(c *fiber.Ctx) error {
		hotelID := c.Params("hotel_id")
		res, err := room.GetRoomByHotelID(c.Context(), hotelID)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.RoomErrorResponse(err))
		}
		return c.JSON(presenter.RoomSuccessResponse(res))
	}
}
