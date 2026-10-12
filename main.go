package main

import (
	"fmt"
	"os"
	"strconv"

	task "github.com/Srinidhi444/CLI-Task-Manager/internal/task"
)

const filename = "data/tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  add <task title> - Add a new task")
		fmt.Println("  list - List all tasks")
		fmt.Println("  complete <task ID> - Mark a task as completed")
		fmt.Println("  delete <task ID> - Delete a task")
		return
	}

	tasks, err := task.LoadTasks(filename)
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . add <task title>")
			return
		}

		title := os.Args[2]

		tasks, err = task.AddTask(tasks, title)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}

		err = task.SaveTasks(filename, tasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}

		fmt.Println("Task added successfully!")

	case "list":
		task.ListTasks(tasks)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . complete <task ID>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil || id <= 0 {
			fmt.Println("Invalid task ID: please enter a positive integer")
			return
		}

		err = task.CompleteTask(tasks, id)
		if err != nil {
			fmt.Printf("Error completing task: %v\n", err)
			return
		}

		err = task.SaveTasks(filename, tasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}

		fmt.Println("Task marked as completed!")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . delete <task ID>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil || id <= 0 {
			fmt.Println("Invalid task ID: please enter a positive integer")
			return
		}

		tasks, err = task.DeleteTask(tasks, id)
		if err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
			return
		}

		err = task.SaveTasks(filename, tasks)
		if err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
			return
		}

		fmt.Println("Task deleted successfully!")

	default:
		fmt.Println("Invalid command. Available commands: add, list, complete, delete")
	}
}
