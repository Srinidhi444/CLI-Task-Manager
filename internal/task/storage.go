package task

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveTasks(filename string, tasks []Task) error {
	res, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %v", err)
	}
	errror := os.WriteFile(filename, res, 0644)
	if errror != nil {
		return fmt.Errorf("failed to write tasks to file: %w", errror)
	}

	return nil
}
func LoadTasks(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("failed to read tasks from file: %w", err)
	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}
	return tasks, nil
}
