package service

import (
	"context"
	"errors"

	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"github.com/prankevich/MyProject/utils"
)

func (s *Service) CreateUser(ctx context.Context, user models.User) (err error) {
	_, err = s.repository.GetUserByName(ctx, user.Username)
	if err != nil {

		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUserNameAlreadyExists
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {

		return err
	}
	if err := s.repository.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *Service) Authenticate(ctx context.Context, user models.User) (int, models.Role, error) {
	userFromDB, err := s.repository.GetUserByName(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return 0, "", errs.ErrUserNotFound
		}

		return 0, "", err
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return 0, "", err
	}

	if userFromDB.Password != user.Password {
		return 0, "", errs.ErrIncorrectUsernameOrPassword
	}

	return userFromDB.ID, userFromDB.Role, nil
}
