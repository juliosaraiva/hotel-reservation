package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/router"
)

const (
	dburi  = "mongodb://localhost:27017"
	dbname = "hotel-reservation"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		return c.Status(code).JSON(fiber.Map{"status_code": code, "error": err.Error()})
	},
}

func main() {
	port := flag.String("port", ":8000", "Port to listen")
	flag.Parse()

	dbconn, cancel, err := db.MongoDBConnection(dburi, dbname)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection sucess!")

	userCollection := dbconn.Collection("users")
	hotelCollection := dbconn.Collection("hotels")
	roomCollection := dbconn.Collection("rooms")

	user := db.NewUser(userCollection)
	hotel := db.NewHotel(hotelCollection)
	room := db.NewRoom(roomCollection, hotel)

	app := fiber.New(config)
	api := app.Group("/api/v1")

	router.UserRouter(api, user)
	router.HotelRouter(api, hotel, room)

	defer cancel()

	fmt.Printf("Starting server on port %v", *port)
	log.Fatal(app.Listen(*port))
}
