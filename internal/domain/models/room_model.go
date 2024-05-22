package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomService struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type        RoomType           `bson:"type" json:"type"`
	Price       float64            `bson:"price" json:"price"`
	HotelID     primitive.ObjectID `bson:"hotel_id" json:"hotel_id"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}

type HotelRooms struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Location string             `bson:"location" json:"location"`
	Rooms    []RoomService      `bson:"rooms" json:"rooms"`
}
