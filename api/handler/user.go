package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v60/github"
	"github.com/juliosaraiva/hotel-reservation/api/db"
	"github.com/juliosaraiva/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
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
	h.userStore.AddNewUser(ctx, &newUser)
	statusCode := fiber.StatusCreated
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusCode,
		"message": "User " + username + " successfully created",
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}
	user, err := h.userStore.GetUserById(ctx, int64(id))
	if err != nil {
		statusCode := fiber.StatusNotFound
		return c.Status(statusCode).JSON(fiber.Map{
			"status":  statusCode,
			"message": "User " + strconv.Itoa(id) + " not found!",
		})
	}

	return c.JSON(user)
}
