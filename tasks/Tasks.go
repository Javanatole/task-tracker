package tasks

import (
	"encoding/json"
	"fmt"
	"time"
)

const DefaultContent = "{\"Tasks\":[]}"

// Tasks representation structure of a filename and his content
type Tasks struct {
	filename string
	content  JSONTasks
}

func deleteElement(slice []Task, index int) []Task {
	return append(slice[:index], slice[index+1:]...)
}

// New constructor of tasks for building new instance of Tasks
func New(filename string) Tasks {
	// first we read the content of the file
	content, err := readContentFromFile(filename)
	if err != nil {
		// in case we discover an error, we re-write the content of the file
		err = writeContentIntoFile(filename, DefaultContent)
		if err != nil {
			// in case we can't re-write the file we launch a panic error
			panic(err)
		}
		// in case
		content = DefaultContent
	}

	contentTasks, err := ParseJSONTasks(content)
	if err != nil {
		contentTasks = JSONTasks{[]Task{}}
	}

	tasks := Tasks{filename: filename, content: contentTasks}

	return tasks
}

func (tasks Tasks) FindLastId() int {
	var maxId = 0
	for _, task := range tasks.content.Tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId
}

func (tasks Tasks) FindIndex(id int) (int, error) {
	for index, task := range tasks.content.Tasks {
		if task.Id == id {
			return index, nil
		}
	}
	return -1, fmt.Errorf("task with id %d not found", id)
}

// AddTask add task to json
func (tasks Tasks) AddTask(description string) {
	newTask := Task{
		Id:          tasks.FindLastId() + 1,
		Description: description,
		Status:      "todo",
		UpdatedAt:   time.Now().String(),
		CreatedAt:   time.Now().String(),
	}

	tasks.content.Tasks = append(tasks.content.Tasks, newTask)
	content, err := json.Marshal(tasks.content)
	if err != nil {
		panic(err)
	}
	err = writeContentIntoFile(tasks.filename, string(content))
	if err != nil {
		panic(err)
	}
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
	content, err := json.Marshal(JSONTasks{Tasks: newTasks})
	if err != nil {
		panic(err)
	}
	err = writeContentIntoFile(tasks.filename, string(content))
	if err != nil {
		panic(err)
	}
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
	content, err := json.Marshal(JSONTasks{Tasks: newTasks})
	if err != nil {
		panic(err)
	}
	err = writeContentIntoFile(tasks.filename, string(content))
	if err != nil {
		panic(err)
	}
}

// DeleteTask delete a task with a id
func (tasks Tasks) DeleteTask(id int) {
	foundTask, err := tasks.FindIndex(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	newTasks := deleteElement(tasks.content.Tasks, foundTask)

	content, err := json.Marshal(JSONTasks{Tasks: newTasks})
	if err != nil {
		panic(err)
	}
	err = writeContentIntoFile(tasks.filename, string(content))
	if err != nil {
		panic(err)
	}
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
