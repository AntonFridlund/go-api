package tasks

import (
	"encoding/json"
	taskModel "main/models/tasks"
	taskService "main/services/tasks"
	"net/http"
	"strconv"
)

// Represents the task controller
type TaskController struct {
	taskService taskService.ITaskService
}

func NewTaskController(taskService taskService.ITaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	data, err := c.taskService.GetAllTasks()
	if err != nil {
		http.Error(w, "Error: failed to retrieve tasks", http.StatusInternalServerError)
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

func (c *TaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.Error(w, "Error: invalid task ID", http.StatusBadRequest)
		return
	}

	data, err := c.taskService.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Error: failed to retrieve task", http.StatusInternalServerError)
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

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task taskModel.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Error: invalid task data"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := task.Validate(); err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.taskService.CreateTask(task); err != nil {
		http.Error(w, "Error: failed to create task", http.StatusInternalServerError)
		return
	}

	taskDTO := taskModel.TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
	}

	jsonData, err := json.Marshal(taskDTO)
	if err != nil {
		http.Error(w, "Error: failed to encode data to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}
