package tasks

import (
	"time"
)

// TaskService manages task operations
type TaskService struct {
	Repo JSONStorage
}

func (taskService *TaskService) AddTask(description string) error {
	tasks, err := taskService.Repo.Load()
	if err != nil {
		return err
	}

	task := Task{
		Id:          tasks.FindLastId() + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	tasks.Tasks = append(tasks.Tasks, task)

	return taskService.Repo.Save(tasks)
}

func (taskService *TaskService) UpdateTaskStatus(id int, status string) error {
	tasks, err := taskService.Repo.Load()
	if err != nil {
		return err
	}

	index, err := tasks.FindIndex(id)
	if err != nil {
		return err
	}

	tasks.Tasks[index].Status = status
	tasks.Tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)

	return taskService.Repo.Save(tasks)
}

func (taskService *TaskService) UpdateTaskDescription(id int, description string) error {
	tasks, err := taskService.Repo.Load()
	if err != nil {
		return err
	}

	index, err := tasks.FindIndex(id)
	if err != nil {
		return err
	}

	tasks.Tasks[index].Description = description
	tasks.Tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)

	return taskService.Repo.Save(tasks)
}

func (taskService *TaskService) DeleteTask(id int) error {
	tasks, err := taskService.Repo.Load()
	if err != nil {
		return err
	}

	index, err := tasks.FindIndex(id)
	if err != nil {
		return err
	}

	tasks.DeleteElement(index)
	return taskService.Repo.Save(tasks)
}

func (taskService *TaskService) ListTasks(status string) ([]Task, error) {
	tasks, err := taskService.Repo.Load()
	if err != nil {
		return nil, err
	}

	if status == "" {
		return tasks.Tasks, nil
	}

	var filtered []Task
	for _, task := range tasks.Tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}
