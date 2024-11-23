package tasks

import (
	taskController "main/controllers/tasks"
	"main/middlewares/auth"
	taskService "main/services/tasks"
	"net/http"
)

// Initialize services and controllers
var tskService = taskService.NewTaskService()
var tskController = taskController.NewTaskController(tskService)

// NewRouter creates the tasks router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/tasks", tskController.GetAllTasks)
	router.HandleFunc("GET /api/v1/tasks/{id}", tskController.GetTaskByID)
	router.HandleFunc("POST /api/v1/tasks/create", auth.AuthMiddleware(tskController.CreateTask))

	return router
}
