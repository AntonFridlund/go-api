package tasks

import taskModel "main/models/tasks"

// Service interface for dependency injection
type ITaskService interface {
	GetAllTasks() ([]taskModel.TaskDTO, error)
	GetTaskByID(id int) (taskModel.TaskDTO, error)
	CreateTask(task taskModel.TaskModel) error
}

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) GetAllTasks() ([]taskModel.TaskDTO, error) {
	tasks := []taskModel.TaskDTO{
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
	return tasks, nil
}

func (s *TaskService) GetTaskByID(id int) (taskModel.TaskDTO, error) {
	task := taskModel.TaskDTO{
		ID:          id,
		Title:       "A specific task",
		Description: "A task which was fetched by ID",
		Done:        true,
	}
	return task, nil
}

func (s *TaskService) CreateTask(task taskModel.TaskModel) error {
	return nil
}
