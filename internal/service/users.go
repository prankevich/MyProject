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

func (s *Service) GetAllUsers() (users []models.User, err error) {
	ctx := context.Background()
	users, err = s.repository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUsersByID(id int) (models.User, error) {
	ctx := context.Background()

	var user models.User
	err := s.cache.Get(ctx, fmt.Sprintf("user:%d", id), &user)
	if err == nil {
		return user, nil // нашли в кэше
	}

	user, err = s.repository.GetUsersByID(ctx, id)
	if err != nil {
		return models.User{}, errs.ErrUserNotfound
	}

	if err = s.cache.Set(ctx, fmt.Sprintf("user:%d", user.ID), user, TTL); err != nil {
		fmt.Printf("error during cache set: %v\n", err)
	}

	return user, nil
}
func (s *Service) CreateUsersByID(users models.User) (err error) {
	ctx := context.Background()
	err = s.repository.CreateUsersByID(ctx, users)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUsersByID(users models.User) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetUsersByID(ctx, users.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}

	err = s.repository.UpdateUsersByID(ctx, users)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteUsersByID(id int) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetUsersByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}

	err = s.repository.DeleteUsersByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
