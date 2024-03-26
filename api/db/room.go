package db

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStorer interface {
	Insert(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	Collection *mongo.Collection
	Hotel      HotelStorer
}

func (s *MongoRoomStore) Insert(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.Collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	_ = s.Hotel.UpdateRooms(ctx, room.HotelID.Hex(), Map{"rooms": room.ID})

	return room, nil
}
