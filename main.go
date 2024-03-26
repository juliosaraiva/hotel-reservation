package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/api/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi    = "mongodb://localhost:27017"
	dbname   = "hotel-reservation"
	userColl = "users"
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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(client)
	}

	userHandler := handler.NewUserHandler(&db.MongoUserStore{
		Collection: client.Database(dbname).Collection("user"),
	})

	hotelStore := db.MongoHotelStore{
		Collection: client.Database(dbname).Collection("hotel"),
	}
	roomStore := db.MongoRoomStore{
		Collection: client.Database(dbname).Collection("room"),
		Hotel:      &hotelStore,
	}

	hotelHandler := handler.NewHotelHandler(
		&hotelStore,
		&roomStore,
	)

	app := fiber.New(config)
	api := app.Group("/api/v1")
	user := api.Group("/user")
	hotel := api.Group("/hotel")

	// User Endpoints
	user.Get("/", userHandler.GetAllUsers)
	user.Get("/:id", userHandler.GetUserById)
	user.Post("/", userHandler.CreateUser)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)

	// Hotel Endpoints
	hotel.Get("/", hotelHandler.GetAll)

	fmt.Printf("Starting server on port %v", *port)
	log.Fatal(app.Listen(*port))
}
