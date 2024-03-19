package handler

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v60/github"
	"github.com/juliosaraiva/hotel-reservation/types"
)

func GetUserByName(c *fiber.Ctx) error {
	ctx := context.Background()
	gh := github.NewClient(nil)

	username := c.Params("username")

	resp, _, err := gh.Users.Get(ctx, username)
	if err != nil {
		log.Fatal("Error fetching user details :%v\n", err)
	}

	user := types.User{
		ID: *resp.ID, Login: *resp.Login, Name: *resp.Name,
	}

	return c.JSON(user)
}
