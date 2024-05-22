package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/interfaces"
	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
	"github.com/juliosaraiva/hotel-reservation/internal/presenter"
)

func AddUser(user interfaces.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.UserParams
		if err := c.BodyParser(&params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := params.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}
		userParams, err := models.NewUserFromParams(params)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		newUser, err := user.AddUser(c.Context(), userParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to create new user"})
		}

		return c.JSON(newUser)
	}
}

func GetUserById(user interfaces.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		user, err := user.GetUserById(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(&user)
	}
}

func GetUsers(user interfaces.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := user.GetUsers(c.Context())
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(users)
	}
}

func UpdateUser(user interfaces.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var params models.UserParams
		if err := c.BodyParser(&params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse"})
		}
		updateUser, err := models.UpdateUserFromParams(params)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unable to update user"})
		}
		if err := user.UpdateUser(c.Context(), id, updateUser); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
	}
}

func DeleteUser(user interfaces.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("id")
		res, err := user.DeleteUser(c.Context(), userId)
		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(presenter.UserResponseError(err))
		}
		return c.JSON(res)
	}
}
