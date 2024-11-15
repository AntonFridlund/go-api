package tasks

import (
	controllerTask "main/controllers/tasks"
	"main/middlewares/auth"
	serviceTask "main/services/tasks"
	"net/http"
)

// Initialize services and controllers
var taskService = serviceTask.NewTaskService()
var taskController = controllerTask.NewTaskController(taskService)

// NewRouter creates the tasks router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/tasks", taskController.GetAllTasks)
	router.HandleFunc("GET /api/v1/tasks/{id}", taskController.GetTaskByID)
	router.HandleFunc("POST /api/v1/tasks/create", auth.AuthMiddleware(taskController.CreateTask))

	return router
}
