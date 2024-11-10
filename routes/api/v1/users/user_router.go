package users

import (
	controllerUser "main/controllers/users"
	"main/middlewares/auth"
	serviceUser "main/services/users"
	"net/http"
)

var userService = serviceUser.NewUserService()
var userController = controllerUser.NewUserController(userService)

// NewRouter creates the users router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/users", userController.GetAllUsers)
	router.HandleFunc("GET /api/v1/users/{id}", userController.GetUserByID)
	router.HandleFunc("POST /api/v1/users/create", auth.AuthMiddleware(userController.CreateUser))

	return router
}
