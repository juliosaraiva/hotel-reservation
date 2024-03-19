package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/handler"
)

func main() {
	port := flag.String("port", ":8000", "Port to listen")
	flag.Parse()

	app := fiber.New()
	api := app.Group("/api/v1")

	api.Get("/user/:username", handler.GetUserByName)

	fmt.Printf("Starting server on port %v", *port)
	app.Listen(*port)
}
