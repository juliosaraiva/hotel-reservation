package main

import (
	"context"
	"fmt"
	"log"

	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBUri  = "mongodb://localhost:27017"
	DBName = "hotel-reservation"
)

var (
	ctx        = context.Background()
	client, _  = mongo.Connect(context.TODO(), options.Client().ApplyURI(DBUri))
	roomStore  db.RoomStorer
	hotelStore db.HotelStorer
)

func seedHotel(name string, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Type:  types.SingleRoomType,
			Price: 100.0,
		},
		{
			Type:  types.DeluxeRoomType,
			Price: 550.5,
		},
		{
			Type:  types.DoubleRoomType,
			Price: 250,
		},
	}

	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		fmt.Println(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.Insert(ctx, &room)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	var err error
	if err = client.Database(DBName).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = &db.MongoHotelStore{
		Collection: client.Database(DBName).Collection("hotel"),
	}

	roomStore = &db.MongoRoomStore{
		Collection: client.Database(DBName).Collection("room"),
		Hotel:      hotelStore,
	}
}
func main() {
	seedHotel("Ibis", "Sao Paulo")
	seedHotel("The cozy hotel", "Las Vegas")
}
