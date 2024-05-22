package db

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type room struct {
	Collection *mongo.Collection
	Hotel      interfaces.Hotel
}

func NewRoom(collection *mongo.Collection, hotel interfaces.Hotel) *room {
	return &room{
		Collection: collection,
		Hotel:      hotel,
	}
}

func (r *room) Insert(ctx context.Context, room *models.Room) (*models.Room, error) {
	resp, err := r.Collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = resp.InsertedID.(primitive.ObjectID)
	_ = r.Hotel.UpdateRooms(ctx, room.HotelID.Hex(), interfaces.Map{"rooms": room.ID})

	return room, nil
}

func (r *room) GetRoom(ctx context.Context, id string) ([]*models.Room, error) {
	var room []*models.Room
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cur, err := r.Collection.Find(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	if err = cur.All(ctx, &room); err != nil {
		return nil, err
	}
	return room, nil
}

func (r *room) GetRooms(ctx context.Context) ([]*models.Room, error) {
	cur, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var rooms []*models.Room
	if err = cur.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *room) GetRoomByHotelID(ctx context.Context, hotelID string) ([]*models.Room, error) {
	oid, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		return nil, err
	}

	cur, err := r.Collection.Find(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}

	var rooms []*models.Room
	if err = cur.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}
