package interfaces

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
)

type Map map[string]interface{}

type Hotel interface {
	Insert(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error)
	GetAll(ctx context.Context) ([]*models.HotelRooms, error)
	Get(ctx context.Context, id string) ([]*models.HotelRooms, error)
	UpdateRooms(context.Context, string, Map) error
}

type Room interface {
	Insert(context.Context, *models.Room) (*models.Room, error)
	GetRoom(ctx context.Context, id string) ([]*models.Room, error)
	GetRooms(ctx context.Context) ([]*models.Room, error)
	GetRoomByHotelID(ctx context.Context, hotelID string) ([]*models.Room, error)
}
