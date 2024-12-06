package auxiliary

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lalith-codeable/cli-task-manager/structs"
)

func ReadFile(filePath string) ([]structs.Task, error) {
	var tasks []structs.Task
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("File %s not found. Creating a new one...\n", filePath)
			dir := strings.Split(filePath, "/")
			fmt.Print(dir)
			os.Mkdir(dir[0], 0777)
			err = os.WriteFile(filePath, []byte("[]"), 0644)
			if err != nil {
				return nil, fmt.Errorf("failed to create file %s: %w", filePath, err)
			}
			return tasks, nil
		}
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON data from file %s: %w", filePath, err)
	}

	return tasks, nil
}

func WriteFile(filePath string, tasks []structs.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks to JSON: %w", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write tasks to file %s: %w", filePath, err)
	}

	return nil
}

func Scream() {
	fmt.Printf(`
Command:

task <subcommand>

Description:	A cli todo manager.

Usage:

task add "<task_name>"			Adds a task to the task list.
task list				Lists all the tasks.
task list done				Lists all tasks marked as done.
task list undone			Lists all tasks marked as undone.
task done <id>				Marks the task as done based on task id.
task undone <id>			Marks the task as undone based on task id.
task delete <id>			Deletes a specific task based on task id.
task reset-all				Deletes all the contents of the tasks table.
		
All Commands:
		
add, done, undone, delete`)
}

func Checks() {
	filePath := "store/taskdata.json"
	dirPath := strings.Split(filePath, "/")[0]
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0777)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.WriteFile(filePath, []byte("[]"), 0644)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
	}
}
