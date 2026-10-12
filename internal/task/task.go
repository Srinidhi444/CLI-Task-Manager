package task

import (
	"fmt"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func AddTask(tasks []Task, title string) ([]Task, error) {
	if strings.TrimSpace(title) == "" {
		return tasks, fmt.Errorf("task title cannot be empty")
	}
	max_id := 0
	for _, task := range tasks {
		if task.ID > max_id {
			max_id = task.ID
		}
	}
	task := Task{
		ID:        max_id + 1,
		Title:     strings.TrimSpace(title),
		Completed: false,
	}
	tasks = append(tasks, task)

	return tasks, nil
}
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Your Tasks:")

	for _, task := range tasks {
		status := "Pending"

		if task.Completed {
			status = "Completed"
		}

		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}
func CompleteTask(tasks []Task, id int) error {
	for index := range tasks {
		if tasks[index].ID == id {
			tasks[index].Completed = true
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}
func DeleteTask(tasks []Task, id int) ([]Task, error) {
	for index := range tasks {
		if tasks[index].ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("task with ID %d not found", id)
}
