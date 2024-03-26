package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/db"
)

type HotelHandler struct {
	HotelStore db.HotelStorer
	RoomStore  db.RoomStorer
}

func NewHotelHandler(h db.HotelStorer, r db.RoomStorer) *HotelHandler {
	return &HotelHandler{
		HotelStore: h,
		RoomStore:  r,
	}
}

func (h *HotelHandler) GetAll(c *fiber.Ctx) error {
	q := c.Queries()

	fmt.Println(q["rooms"])
	hotels, err := h.HotelStore.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(hotels)
}
