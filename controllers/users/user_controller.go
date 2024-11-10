package users

import (
	"encoding/json"
	modelUser "main/models/users"
	serviceUser "main/services/users"
	"net/http"
	"strconv"
)

type userController struct {
	userService serviceUser.IUserService
}

func NewUserController(userService serviceUser.IUserService) *userController {
	return &userController{userService: userService}
}

func (c *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	data, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error: Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error: failed to encode data to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (c *userController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(w, "Error: invalid user ID", http.StatusBadRequest)
		return
	}

	data, err := c.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "Error: failed to retrieve user", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error: failed to encode data to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userModel modelUser.User
	if err := json.NewDecoder(r.Body).Decode(&userModel); err != nil {
		http.Error(w, "Error: invalid user data", http.StatusBadRequest)
		return
	}

	userModel.ID = 1234567890

	if err := userModel.Validate(); err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.userService.CreateUser(userModel); err != nil {
		http.Error(w, "Error: failed to create user", http.StatusInternalServerError)
		return
	}

	userDTO := modelUser.UserDTO{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
	}

	jsonData, err := json.Marshal(userDTO)
	if err != nil {
		http.Error(w, "Error: failed to encode data to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}
