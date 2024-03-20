package handler

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v60/github"
	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/types"
)

type UserHandler struct {
	UserStore db.UserStore
}

func (h *UserHandler) AddNewUser(c *fiber.Ctx) error {
	ctx := context.Background()
	gh := github.NewClient(nil)
	username := c.Params("username")
	user, _, err := gh.Users.Get(ctx, username)

	if err != nil {
		log.Fatal(err)
	}
	newUser := types.User{
		ID: user.ID, Login: user.Login, Name: user.Name,
	}
	h.UserStore.AddNewUser(&newUser)
	statusCode := fiber.StatusCreated
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": "User " + username + "successfully created",
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.UserStore.GetUserById(id)
	if err != nil {
		statusCode := fiber.StatusNotFound
		return c.Status(statusCode).JSON(fiber.Map{
			"status":  statusCode,
			"message": "User " + id + " not found!",
		})
	}

	return c.JSON(user)
}
