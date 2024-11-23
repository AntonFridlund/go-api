package users

import userModel "main/models/users"

// Service interface for dependency injection
type IUserService interface {
	GetAllUsers() ([]userModel.UserDTO, error)
	GetUserByID(id int) (userModel.UserDTO, error)
	CreateUser(user userModel.UserModel) error
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() ([]userModel.UserDTO, error) {
	users := []userModel.UserDTO{
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
	return users, nil
}

func (s *UserService) GetUserByID(id int) (userModel.UserDTO, error) {
	user := userModel.UserDTO{
		ID:        id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	return user, nil
}

func (s *UserService) CreateUser(user userModel.UserModel) error {
	return nil
}
