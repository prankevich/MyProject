package service

import (
	"context"
	"errors"
	"github.com/prankevich/MyProject/internal/config"
	"github.com/prankevich/MyProject/internal/errs"
	"github.com/prankevich/MyProject/internal/models"
	"github.com/prankevich/MyProject/pkg"
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

func (s *Service) Authenticate(ctx context.Context, user models.User) (string, error) {

	userFromDB, err := s.repository.GetUserByName(ctx, user.Username)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return "", errs.ErrUserNotFound
		}

		return "", err
	}

	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return "", err
	}
	if userFromDB.Password != user.Password {
		return "", errs.ErrIncorrectUsernameOrPassword
	}
	token, err := pkg.GenerateToken(userFromDB.ID, config.AppSettings.AuthParams.TtlMinutes)
	if err != nil {
		return "", err
	}
	return token, nil
}
