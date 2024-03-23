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
		Client:     client,
		Collection: client.Database(dbname).Collection(userColl),
	})

	app := fiber.New(config)
	api := app.Group("/api/v1")

	api.Get("/user", userHandler.GetAllUsers)
	api.Get("/user/:id", userHandler.GetUserById)
	api.Post("/user", userHandler.CreateUser)
	api.Put("user/:id", userHandler.UpdateUser)
	api.Delete("/user/:id", userHandler.DeleteUser)

	fmt.Printf("Starting server on port %v", *port)
	log.Fatal(app.Listen(*port))
}
