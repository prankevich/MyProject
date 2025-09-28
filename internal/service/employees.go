package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"time"
)

var (
	TTL = time.Minute * 3
)

func (s *Service) GetAllEmployees() (employees []models.Employees, err error) {
	ctx := context.Background()
	employees, err = s.repository.GetAllEmployees(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *Service) GetEmployeesByID(id int) (models.Employees, error) {
	ctx := context.Background()

	var employees models.Employees
	err := s.cache.Get(ctx, fmt.Sprintf("user:%d", id), &employees)
	if err == nil {

		return employees, nil // нашли в кэше
	}

	employees, err = s.repository.GetEmployeesByID(ctx, id)
	if err != nil {
		return models.Employees{}, errs.ErrEmployeesNotfound
	}

	if err = s.cache.Set(ctx, fmt.Sprintf("user:%d", employees.ID), employees, TTL); err != nil {
		fmt.Printf("error during cache set: %v\n", err)
	}

	return employees, nil
}
func (s *Service) CreateEmployees(users models.Employees) (err error) {
	ctx := context.Background()
	err = s.repository.CreateEmployees(ctx, users)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateEmployeesByID(employees models.Employees) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetEmployeesByID(ctx, employees.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrEmployeesNotfound
		}
		return err
	}

	err = s.repository.UpdateEmployeesByID(ctx, employees)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteEmployeesByID(id int) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetEmployeesByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrEmployeesNotfound
		}
		return err
	}

	err = s.repository.DeleteEmployeesByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
