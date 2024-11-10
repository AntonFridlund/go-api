package users

import modelUser "main/models/users"

type IUserService interface {
	GetAllUsers() ([]modelUser.UserDTO, error)
	GetUserByID(id int) (modelUser.UserDTO, error)
	CreateUser(user modelUser.User) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() ([]modelUser.UserDTO, error) {
	userModels := []modelUser.UserDTO{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@example.com",
		},
	}
	return userModels, nil
}

func (s *UserService) GetUserByID(id int) (modelUser.UserDTO, error) {
	userModel := modelUser.UserDTO{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	return userModel, nil
}

func (s *UserService) CreateUser(user modelUser.User) error {
	return nil
}
