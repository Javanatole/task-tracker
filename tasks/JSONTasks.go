package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Task One unitary task
type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
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

// JSONTasks representation of all tasks
type JSONTasks struct {
	Tasks []Task
}

// String representation of all tasks
func (jsonTasks JSONTasks) String() string {
	var buffer bytes.Buffer
	for _, task := range jsonTasks.Tasks {
		buffer.WriteString(task.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// DeleteElement delete element from list of tasks
func (jsonTasks JSONTasks) DeleteElement(index int) []Task {
	return append(jsonTasks.Tasks[:index], jsonTasks.Tasks[index+1:]...)
}

// FindLastId find last id of tasks
func (jsonTasks JSONTasks) FindLastId() int {
	var maxId = 0
	for _, task := range jsonTasks.Tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId
}

// FindIndex find index of specific id
func (jsonTasks JSONTasks) FindIndex(id int) (int, error) {
	for index, task := range jsonTasks.Tasks {
		if task.Id == id {
			return index, nil
		}
	}
	return -1, fmt.Errorf("task with id %d not found", id)
}

// ParseJSONTasks generate json tasks from string
func ParseJSONTasks(content string) (JSONTasks, error) {
	var tasks JSONTasks
	err := json.Unmarshal([]byte(content), &tasks)
	return tasks, err
}