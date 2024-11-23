package tasks

import "errors"

type TaskModel struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Model validation logic
func (t *TaskModel) Validate() error {
	if len(t.Title) < 1 {
		return errors.New("empty task title")
	} else if len(t.Description) < 1 {
		return errors.New("empty task description")
	}
	return nil
}
