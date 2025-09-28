package contracts

import (
	"context"
	"github.com/prankevich/MyProject/internal/models"
)

type RepositoryI interface {
	GetAllEmployees(ctx context.Context) (employees []models.Employees, err error)
	GetEmployeesByID(ctx context.Context, id int) (employees models.Employees, err error)
	CreateEmployees(ctx context.Context, employees models.Employees) (err error)
	UpdateEmployeesByID(ctx context.Context, employees models.Employees) (err error)
	DeleteEmployeesByID(ctx context.Context, id int) (err error)
	GetUserByName(ctx context.Context, username string) (user models.User, err error)
	CreateUser(ctx context.Context, users models.User) (err error)
	GetUserByID(ctx context.Context, username string) (user models.User, err error)
}
