package handler

import (
	"github.com/gofiber/fiber/v2"
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

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		statusCode := fiber.StatusNotFound
		return c.Status(statusCode).JSON(fiber.Map{"status": statusCode, "msg": err.Error()})
	}
	return c.JSON(&user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUser(c.Context())
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var params types.UserParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Line 41"})
	}

	if err := params.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validate"})
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Line 48"})
	}

	newUser, err := h.userStore.AddUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Line 53"})
	}

	return c.JSON(newUser)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var params types.UserParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse"})
	}
	updateUser, err := types.UpdateUserFromParams(params)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unable to update user"})
	}
	if err := h.userStore.UpdateUser(c.Context(), id, updateUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	res, err := h.userStore.DeleteUser(c.Context(), userId)
	statusCode := fiber.StatusNoContent
	if err != nil {
		return c.Status(statusCode).JSON(fiber.Map{"msg": err})
	}
	return c.Status(statusCode).JSON(fiber.Map{"deleted": res.DeletedCount})
}
