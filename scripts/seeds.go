package main

import (
	"context"
	"fmt"
	"log"

	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DBUri  = "mongodb://localhost:27017"
	DBName = "hotel-reservation"
)

var (
	ctx   = context.Background()
	room  interfaces.Room
	hotel interfaces.Hotel
)

func seedHotel(name string, location string) {
	h := models.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []models.Room{
		{
			Type:  models.SingleRoomType,
			Price: 100.0,
		},
		{
			Type:  models.DeluxeRoomType,
			Price: 550.5,
		},
		{
			Type:  models.DoubleRoomType,
			Price: 250,
		},
	}

	insertedHotel, err := hotel.Insert(ctx, &h)
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range rooms {
		r.HotelID = insertedHotel.ID
		_, err := room.Insert(ctx, &r)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func init() {
	var err error
	dbconn, cancel, err := db.MongoDBConnection(DBUri, DBName)
	if err != nil {
		log.Fatal(err)
	}
	hotelCollection := dbconn.Collection("hotels")
	roomCollection := dbconn.Collection("rooms")

	hotel = db.NewHotel(hotelCollection)
	room = db.NewRoom(roomCollection, hotel)

	defer cancel()
}

func main() {
	seedHotel("Ibis", "Sao Paulo")
	seedHotel("The cozy hotel", "Las Vegas")
}
