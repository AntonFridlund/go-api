package v1

import (
	"main/routes/api/v1/tasks"
	"main/routes/api/v1/users"
	"net/http"
)

// NewRouter creates the v1 API router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	userRouter := users.NewRouter()
	router.Handle("/api/v1/users", userRouter)
	router.Handle("/api/v1/users/", userRouter)

	taskRouter := tasks.NewRouter()
	router.Handle("/api/v1/tasks", taskRouter)
	router.Handle("/api/v1/tasks/", taskRouter)

	return router
}
