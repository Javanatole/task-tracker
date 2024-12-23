package main

import (
	"cli-task/tasks"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: go run main.go command options")
		return
	}

	// File name to store the tasks
	filename := "tasks.json"
	// Default content if the file is empty or doesn't exist
	defaultContent := "{\"tasks\":[]}"

	// Create storage and repository
	fileStorage := &tasks.FileTaskStorage{Filename: filename}
	jsonRepository := &tasks.JSONTaskRepository{
		Storage:        fileStorage,
		DefaultContent: defaultContent,
	}

	// Initialize the TaskService
	taskService := &tasks.TaskService{
		Repo: jsonRepository,
	}

	switch os.Args[1] {
	case "add":
		if os.Args[2] == "" {
			fmt.Println("usage: go run main.go add tasks")
			fmt.Println("missing parameters")
		} else {
			// call add function in database
			err := taskService.AddTask(os.Args[2])
			if err != nil {
				panic(err)
			}

		}
	case "delete":
		if os.Args[2] == "" {
			fmt.Println("usage: go run main.go delete id")
			fmt.Println("missing parameters")
		} else {
			// call add function in database
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("ID must be an integer")
			}
			err = taskService.DeleteTask(id)
			if err != nil {
				panic(err)
			}
		}
	case "mark-in-progress":
		if os.Args[2] == "" {
			fmt.Println("usage: go run main.go delete id")
			fmt.Println("missing parameters")
		} else {
			// call add function in database
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("ID must be an integer")
			}
			err = taskService.UpdateTaskStatus(id, "in-progress")
			if err != nil {
				panic(err)
			}
		}
	case "mark-done":
		if os.Args[2] == "" {
			fmt.Println("usage: go run main.go delete id")
			fmt.Println("missing parameters")
		} else {
			// call add function in database
			id, err := strconv.Atoi(os.Args[2])
			err = taskService.UpdateTaskStatus(id, "done")
			if err != nil {
				panic(err)
			}
		}
	case "list":
		filter := ""
		if len(os.Args) == 3 {
			filter = os.Args[2]
		}
		listTasks, err := taskService.ListTasks(filter)
		if err != nil {
			panic(err)
		}
		for _, task := range listTasks {
			fmt.Println(task)
		}
	case "update":
		if len(os.Args) != 4 {
			fmt.Println("usage: go run main.go update id newDescription")
		} else {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("ID must be an integer")
			}
			err = taskService.UpdateTaskDescription(id, os.Args[3])
			if err != nil {
				panic(err)
			}
		}
	default:
		fmt.Println("usage: go run main.go add \"your task\"")
	}
}
