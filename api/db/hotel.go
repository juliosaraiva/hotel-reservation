package db

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStorer interface {
	Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
	GetAll(ctx context.Context) ([]*types.Hotel, error)
	UpdateRooms(context.Context, string, Map) error
}

type MongoHotelStore struct {
	Collection *mongo.Collection
}

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := s.Collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) UpdateRooms(ctx context.Context, id string, params Map) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mapHotel, err := bson.Marshal(bson.M{"$push": params})
	if err != nil {
		return err
	}

	_, err = s.Collection.UpdateOne(ctx, bson.M{"_id": oid}, mapHotel)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoHotelStore) GetAll(ctx context.Context) ([]*types.Hotel, error) {
	var hotel []*types.Hotel
	cur, err := s.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &hotel); err != nil {
		return nil, err
	}

	return hotel, nil
}
