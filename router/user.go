package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/api/handler"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
)

func UserRouter(api fiber.Router, u interfaces.User) {
	user := api.Group("/user")

	user.Get("/", handler.GetUsers(u))
	user.Get("/:id", handler.GetUserById(u))
	user.Post("/", handler.AddUser(u))
	user.Put("/:id", handler.UpdateUser(u))
	user.Delete("/:id", handler.DeleteUser(u))
}
