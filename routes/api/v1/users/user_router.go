package users

import (
	userController "main/controllers/users"
	"main/middlewares/auth"
	userService "main/services/users"
	"net/http"
)

// Initialize services and controllers
var usrService = userService.NewUserService()
var usrController = userController.NewUserController(usrService)

// NewRouter creates the users router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/users", usrController.GetAllUsers)
	router.HandleFunc("GET /api/v1/users/{id}", usrController.GetUserByID)
	router.HandleFunc("POST /api/v1/users/create", auth.AuthMiddleware(usrController.CreateUser))

	return router
}
