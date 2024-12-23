package tasks

import (
	"bytes"
	"errors"
	"fmt"
)

// Task represents a single task
type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// String show task representation
func (task Task) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ID: %d\n", task.Id))
	buffer.WriteString(fmt.Sprintf("Description: %s\n", task.Description))
	buffer.WriteString(fmt.Sprintf("Status: %s\n", task.Status))
	buffer.WriteString(fmt.Sprintf("Created: %s\n", task.CreatedAt))
	buffer.WriteString(fmt.Sprintf("Updated: %s\n", task.UpdatedAt))
	return buffer.String()
}

// JSONTasks represents a collection of tasks
type JSONTasks struct {
	Tasks []Task `json:"tasks"`
}

// String representation of all tasks
func (jsonTasks *JSONTasks) String() string {
	var buffer bytes.Buffer
	for _, task := range jsonTasks.Tasks {
		buffer.WriteString(task.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (jsonTasks *JSONTasks) FindLastId() int {
	maxId := 0
	for _, task := range jsonTasks.Tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId
}

func (jsonTasks *JSONTasks) FindIndex(id int) (int, error) {
	for i, task := range jsonTasks.Tasks {
		if task.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("task not found")
}

func (jsonTasks *JSONTasks) DeleteElement(index int) {
	jsonTasks.Tasks = append(jsonTasks.Tasks[:index], jsonTasks.Tasks[index+1:]...)
}
