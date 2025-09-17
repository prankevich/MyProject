package repository

import "github.com/prankevich/MyProject/models"

func (r *Repository) GetAllUsers() (users []models.User, err error) {
	if err = r.db.Select(&users, `
		SELECT id, name, email, age
		FROM users
		ORDER BY id`); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUsersByID(id int) (user models.User, err error) {
	if err = r.db.Get(&user, `
		SELECT id, name, email, age
		FROM users
		WHERE id = $1`, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) CreateUsers(users models.User) (err error) {
	_, err = r.db.Exec(`INSERT INTO users (name, email, age)
					VALUES ($1, $2, $3)`,
		users.Name,
		users.Email,
		users.Age)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUsersByID(users models.User) (err error) {
	_, err = r.db.Exec(`
		UPDATE users SET name = $1, 
		                    email = $2, 
		                    age = $3
		                    		                WHERE id = $4`,
		users.Name,
		users.Email,
		users.Age,
		users.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUsersByID(id int) (err error) {
	_, err = r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
