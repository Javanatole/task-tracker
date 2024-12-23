package tasks

import (
	"fmt"
	"time"
)

const DefaultContent = "{\"Tasks\":[]}"

// Tasks representation structure of a filename and his content
type Tasks struct {
	jsonHelper JsonHelper
	content    JSONTasks
}

// New constructor of tasks for building new instance of Tasks
func New(filename string) Tasks {
	fileHelper := FileHelper{filename: filename}
	jsonHelper := JsonHelper{fileHelper: fileHelper}

	return Tasks{jsonHelper: jsonHelper, content: jsonHelper.readContentFromFile()}
}

// AddTask add task to json
func (tasks Tasks) AddTask(description string) {
	newTask := Task{
		Id:          tasks.content.FindLastId() + 1,
		Description: description,
		Status:      "todo",
		UpdatedAt:   time.Now().String(),
		CreatedAt:   time.Now().String(),
	}
	tasks.content.Tasks = append(tasks.content.Tasks, newTask)
	tasks.jsonHelper.writeJSONTasks(tasks.content)
}

func (tasks Tasks) MarkAs(id int, status string) {
	var newTasks []Task
	for _, task := range tasks.content.Tasks {
		if task.Id == id {
			task.Status = status
			task.UpdatedAt = time.Now().String()
		}
		newTasks = append(newTasks, task)
	}
	tasks.jsonHelper.writeJSONTasks(JSONTasks{Tasks: newTasks})
}

// UpdateDescriptionTask update description
func (tasks Tasks) UpdateDescriptionTask(id int, description string) {
	var newTasks []Task
	for _, task := range tasks.content.Tasks {
		if task.Id == id {
			task.Description = description
			task.UpdatedAt = time.Now().String()
		}
		newTasks = append(newTasks, task)
	}
	tasks.jsonHelper.writeJSONTasks(JSONTasks{Tasks: newTasks})
}

// DeleteTask delete a task with a id
func (tasks Tasks) DeleteTask(id int) {
	foundTask, err := tasks.content.FindIndex(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	newTasks := tasks.content.DeleteElement(foundTask)
	tasks.jsonHelper.writeJSONTasks(JSONTasks{Tasks: newTasks})
}

// ListTasks list tasks with the given status
// Do not filter if status is equals to empty string
func (tasks Tasks) ListTasks(status string) {
	if status == "" {
		fmt.Println(tasks.content)
		return
	} else {
		var filteredTasks []Task
		for _, task := range tasks.content.Tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
		fmt.Println(JSONTasks{Tasks: filteredTasks})
		return
	}
}
