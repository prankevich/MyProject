package contracts

import (
	"github.com/prankevich/MyProject/internal/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllUsers() (users []models.User, err error)
	GetUsersByID(id int) (product models.User, err error)
	CreateUsersByID(users models.User) (err error)
	UpdateUsersByID(users models.User) (err error)
	DeleteUsersByID(id int) (err error)
}
