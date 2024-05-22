package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/handler"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
)

func HotelRouter(api fiber.Router, h interfaces.Hotel, r interfaces.Room) {
	// Hotels
	hotel := api.Group("/hotel")
	hotel.Get("/", handler.GetHotels(h))
	hotel.Get("/:id", handler.GetHotelDetails(h))

	// Rooms
	rooms := api.Group("/rooms")
	rooms.Get("/", handler.GetRooms(r))
	rooms.Get("/:id", handler.GetRoom(r))
	rooms.Get("/hotel/:hotel_id", handler.GetRoomsByHotelID(r))
}
