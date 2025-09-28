package repository

import (
	"context"
	"github.com/prankevich/MyProject/internal/models"
)

func (r *Repository) GetUserByID(ctx context.Context, username string) (user models.User, err error) {
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, full_name, user_name, password, create_at, update_at 
		FROM users
		WHERE id = $1`, username); err != nil {
		r.logger.Error().Err(err).Str("func", "repository.GetUserByID").Msg("Error selecting users")
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO users (full_name, user_name, password)
					VALUES ($1, $2, $3)`,
		user.FullName,
		user.Username,
		user.Password)
	if err != nil {
		r.logger.Error().Err(err).Str("func", "repository.CreateUser").Msg("Error inserting users")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetUserByName(ctx context.Context, username string) (user models.User, err error) {
	if err = r.db.GetContext(ctx, &user, `
		SELECT full_name, user_name, password, create_at, update_at 
		FROM users
		WHERE user_name = $1`, username); err != nil {
		r.logger.Error().Err(err).Str("func", "repository.GetUserByName").Msg("Error selecting users")
		return models.User{}, r.translateError(err)
	}

	return user, nil
}
