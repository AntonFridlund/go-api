package users

import (
	"encoding/json"
	userModel "main/models/users"
	userService "main/services/users"
	"net/http"
	"strconv"
)

// Represents the user controller
type userController struct {
	userService userService.IUserService
}

func NewUserController(userService userService.IUserService) *userController {
	return &userController{userService: userService}
}

func (c *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	data, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error: failed to retrieve users", http.StatusInternalServerError)
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
	var user userModel.UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error: invalid user data", http.StatusBadRequest)
		return
	}

	user.ID = 1234567890

	if err := user.Validate(); err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.userService.CreateUser(user); err != nil {
		http.Error(w, "Error: failed to create user", http.StatusInternalServerError)
		return
	}

	userDTO := userModel.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
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
