package db

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type hotel struct {
	Collection *mongo.Collection
}

func NewHotel(collection *mongo.Collection) *hotel {
	return &hotel{
		Collection: collection,
	}
}

func (h *hotel) Insert(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	resp, err := h.Collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (h *hotel) UpdateRooms(ctx context.Context, id string, params interfaces.Map) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	mapHotel, err := bson.Marshal(bson.M{"$push": params})
	if err != nil {
		return err
	}
	_, err = h.Collection.UpdateOne(ctx, bson.M{"_id": oid}, mapHotel)
	if err != nil {
		return err
	}

	return nil
}

func (h *hotel) Get(ctx context.Context, id string) ([]*models.HotelRooms, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"_id", oid}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "rooms"},
			{"localField", "_id"},
			{"foreignField", "hotel_id"},
			{"as", "rooms"},
		}}},
	}
	cur, err := h.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var hotel []*models.HotelRooms
	if err = cur.All(ctx, &hotel); err != nil {
		return nil, err
	}

	if len(hotel) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return hotel, nil
}

func (h *hotel) GetAll(ctx context.Context) ([]*models.HotelRooms, error) {
	pipeline := mongo.Pipeline{
		bson.D{{"$lookup", bson.D{
			{"from", "rooms"},
			{"localField", "_id"},
			{"foreignField", "hotel_id"},
			{"as", "rooms"},
		}}},
	}

	cur, err := h.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var hotel []*models.HotelRooms
	if err = cur.All(ctx, &hotel); err != nil {
		return nil, err
	}
	return hotel, nil
}
