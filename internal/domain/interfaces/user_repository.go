package interfaces

import (
	"context"

	"github.com/juliosaraiva/hotel-reservation/internal/domain/models"
)

type Dropper interface {
	Drop(ctx context.Context) error
}

type User interface {
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	AddUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, id string, params *models.UserUpdate) error
	DeleteUser(ctx context.Context, id string) (*models.DeleteResult, error)
	Dropper
}
