package main

import (
	tasks2 "cli-task/tasks"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: go run main.go command options")
		return
	}

	tasks := tasks2.New("tasks.json")

	switch os.Args[1] {
	case "add":
		if os.Args[2] == "" {
			fmt.Println("usage: go run main.go add tasks")
			fmt.Println("missing parameters")
		} else {
			// call add function in database
			tasks.AddTask(os.Args[2])

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
			tasks.DeleteTask(id)
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
			tasks.MarkAs(id, "in-progress")
		}
	case "mark-done":
		if os.Args[2] == "" {
			fmt.Println("usage: go run main.go delete id")
			fmt.Println("missing parameters")
		} else {
			// call add function in database
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("ID must be an integer")
			}
			tasks.MarkAs(id, "done")
		}
	case "list":
		if len(os.Args) == 2 {
			tasks.ListTasks("")
		} else {
			tasks.ListTasks(os.Args[2])
		}
	case "update":
		if len(os.Args) != 4 {
			fmt.Println("usage: go run main.go update id newDescription")
		} else {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("ID must be an integer")
			}
			tasks.UpdateDescriptionTask(id, os.Args[3])
		}
	default:
		fmt.Println("usage: go run main.go add \"your task\"")
	}
}
