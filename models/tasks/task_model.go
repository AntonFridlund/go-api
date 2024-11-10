package tasks

import "errors"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (t *Task) Validate() error {
	if len(t.Title) < 1 {
		return errors.New("empty task title")
	} else if len(t.Description) < 1 {
		return errors.New("empty task description")
	}
	return nil
}
