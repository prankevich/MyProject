package service

import "github.com/prankevich/MyProject/models"

func (s *Service) GetAllUsers() (users []models.User, err error) {
	users, err = s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUsersByID(id int) (users models.User, err error) {
	users, err = s.repository.GetUsersByID(id)
	if err != nil {
		return models.User{}, err
	}

	return users, nil
}

func (s *Service) CreateUsers(users models.User) (err error) {
	err = s.repository.CreateUsers(users)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUsersByID(users models.User) (err error) {
	err = s.repository.UpdateUsersByID(users)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUsersByID(id int) (err error) {
	err = s.repository.DeleteUsersByID(id)
	if err != nil {
		return err
	}

	return nil
}
