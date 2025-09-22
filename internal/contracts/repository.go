package contracts

import (
	"context"
	"github.com/prankevich/MyProject/internal/models"
)

type RepositoryI interface {
	GetAllUsers(ctx context.Context) (products []models.User, err error)
	GetUsersByID(ctx context.Context, id int) (product models.User, err error)
	CreateUsersByID(ctx context.Context, product models.User) (err error)
	UpdateUsersByID(ctx context.Context, product models.User) (err error)
	DeleteUsersByID(ctx context.Context, id int) (err error)
}
