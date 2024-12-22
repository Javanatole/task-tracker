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

func (task Task) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ID: %d\n", task.Id))
	buffer.WriteString(fmt.Sprintf("Description: %s\n", task.Description))
	buffer.WriteString(fmt.Sprintf("Status: %s\n", task.Status))
	buffer.WriteString(fmt.Sprintf("Created: %s\n", task.CreatedAt))
	buffer.WriteString(fmt.Sprintf("Updated: %s\n", task.UpdatedAt))
	return buffer.String()
}

// JSONTasks JSON representation of all tasks
type JSONTasks struct {
	Tasks []Task
}

func (jsonTasks JSONTasks) String() string {
	var buffer bytes.Buffer
	for _, task := range jsonTasks.Tasks {
		buffer.WriteString(task.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// ParseJSONTasks generate json tasks from string
func ParseJSONTasks(content string) (JSONTasks, error) {
	var tasks JSONTasks
	err := json.Unmarshal([]byte(content), &tasks)
	return tasks, err
}
