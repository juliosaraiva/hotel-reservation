package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
}

type HotelUpdate struct {
	Name     string               `bson:"name,omitempty" json:"name,omitempty"`
	Location string               `bson:"location,omitempty" json:"location,omitempty"`
	Rooms    []primitive.ObjectID `bson:"rooms,omitempty" json:"rooms,omitempty"`
}

const (
	SingleRoomType RoomType = iota
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type RoomType int

type Room struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type    RoomType           `bson:"type" json:"type"`
	Price   float64            `bson:"price" json:"price"`
	HotelID primitive.ObjectID `bson:"hotel_id" json:"hotel_id"`
}
