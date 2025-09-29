package contracts

import (
	"context"
	"github.com/prankevich/MyProject/internal/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceI interface {
	GetAllEmployees() (users []models.Employees, err error)
	GetEmployeesByID(id int) (product models.Employees, err error)
	CreateEmployees(users models.Employees) (err error)
	UpdateEmployeesByID(users models.Employees) (err error)
	DeleteEmployeesByID(id int) (err error)
	CreateUser(ctx context.Context, users models.User) (err error)
	Authenticate(ctx context.Context, user models.User) (int, models.Role, error)
}
