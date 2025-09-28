package repository

import (
	"context"
	"github.com/prankevich/MyProject/internal/models"
)

func (r *Repository) GetAllEmployees(ctx context.Context) (employees []models.Employees, err error) {
	if err = r.db.SelectContext(ctx, &employees, `
		SELECT id, name, email, age
		FROM employees
		ORDER BY id`); err != nil {
		r.logger.Error().Err(err).Str("func", "repository.GetAllEmployees").Msg("Error selecting users")

	}

	return employees, nil
}

func (r *Repository) GetEmployeesByID(ctx context.Context, id int) (employees models.Employees, err error) {
	if err = r.db.GetContext(ctx, &employees, `
		SELECT id, name, email, age
		FROM employees
		WHERE id = $1`, id); err != nil {
		r.logger.Error().Err(err).Str("func", "repository.GetEmployeesByID").Msg("Error selecting users")
		return models.Employees{}, err
	}

	return employees, nil
}

func (r *Repository) CreateEmployees(ctx context.Context, employees models.Employees) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO employees (name, email, age)
					VALUES ($1, $2, $3)`,
		employees.Name,
		employees.Email,
		employees.Age)
	if err != nil {
		r.logger.Error().Err(err).Str("func", "repository.CreateEmployees").Msg("Error inserting employees")
		return err
	}

	return nil
}

func (r *Repository) UpdateEmployeesByID(ctx context.Context, employees models.Employees) (err error) {
	_, err = r.db.ExecContext(ctx, `
		UPDATE employees SET name = $1, 
		                    email = $2, 
		                    age = $3
		                    		                WHERE id = $4`,
		employees.Name,
		employees.Email,
		employees.Age,
		employees.ID)
	if err != nil {

		r.logger.Error().Err(err).Str("func", "repository.UpdateEmployeesByID").Msg("Error updating employee ")

		return r.translateError(err)
	}

	return nil
}

func (r *Repository) DeleteEmployeesByID(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, `DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		r.logger.Error().Err(err).Int("Employees id", id).Str("func", "repository.DeleteEmployeesByID").Msg("Error deleting  employees")
		return err
	}

	return nil
}
