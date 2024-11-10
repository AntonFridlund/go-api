package tasks

import modelTask "main/models/tasks"

type ITaskService interface {
	GetAllTasks() ([]modelTask.TaskDTO, error)
	GetTaskByID(id int) (modelTask.TaskDTO, error)
	CreateTask(task modelTask.Task) error
}

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) GetAllTasks() ([]modelTask.TaskDTO, error) {
	taskModels := []modelTask.TaskDTO{
		{
			ID:          1,
			Title:       "Do this first",
			Description: "A task which should be done first",
			Done:        false,
		},
		{
			ID:          2,
			Title:       "Do this second",
			Description: "A task which should be done second",
			Done:        false,
		},
	}
	return taskModels, nil
}

func (s *TaskService) GetTaskByID(id int) (modelTask.TaskDTO, error) {
	taskModel := modelTask.TaskDTO{
		ID:          id,
		Title:       "A specific task",
		Description: "A task which was fetched by ID",
		Done:        true,
	}
	return taskModel, nil
}

func (s *TaskService) CreateTask(task modelTask.Task) error {
	return nil
}
