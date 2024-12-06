package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lalith-codeable/cli-task-manager/auxiliary"
	"github.com/lalith-codeable/cli-task-manager/structs"
	"github.com/olekukonko/tablewriter"
)

const filePath = "store/taskdata.json"

func Add(name string) {

	tasks, err := auxiliary.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	t := structs.Task{
		Id:         uuid.New(),
		Name:       name,
		Completed:  false,
		Created_at: time.Now(),
	}

	tasks = append(tasks, t)

	err = auxiliary.WriteFile(filePath, tasks)
	if err != nil {
		fmt.Printf("Error writing tasks: %v\n", err)
		return
	}

	fmt.Printf("New task added: %s\n", t.Name)
}

func Toggle(index uint8) {
	tasks, err := auxiliary.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	if int(index) >= len(tasks) {
		fmt.Printf("\033[31mInvalid task id %v - Action failed.\033[0m\n", index)
		return
	}
	tasks[index].Completed = !tasks[index].Completed
	werr := auxiliary.WriteFile(filePath, tasks)
	if werr != nil {
		fmt.Printf("Error updating tasks: %v\n", werr)
	}
	fmt.Printf("Task:%s Completed:%v\n", tasks[index].Name, tasks[index].Completed)
}

func Delete(ids []uint8) {
	tasks, err := auxiliary.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	if len(ids) > 0 {

		toDelete := make(map[uint8]bool)
		for _, id := range ids {
			toDelete[id] = true
		}

		var updatedTasks []structs.Task
		for i, task := range tasks {
			if !toDelete[uint8(i)] {
				updatedTasks = append(updatedTasks, task)
			}
		}
		fmt.Printf("Deleted tasks: %v\n", ids)

		err := auxiliary.WriteFile(filePath, updatedTasks)
		if err != nil {
			fmt.Printf("Error writing updated tasks: %v\n", err)
		}
	} else {
		fmt.Println("Deleted tasks: None.")
	}
}

func List() {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("\033[31mError reading tasks: %v\033[0m\n", err)
		return
	}

	var tasks []structs.Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Printf("\033[31mError parsing tasks: %v\033[0m\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Completed", "Created At"})

	for slno, task := range tasks {

		status := "\033[31m✗\033[0m"

		if task.Completed {
			status = "\033[32m✔\033[0m"
		}
		createdAt := task.Created_at.Format("02 Jan 2006, 03:04 PM")
		table.Append([]string{
			strconv.Itoa(slno),
			task.Name,
			status,
			createdAt,
		})

		table.SetBorder(true)
		table.SetRowLine(true)
		table.SetAlignment(tablewriter.ALIGN_CENTER)

	}
	table.Render()
}
func ListDone() {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("\033[31mError reading tasks: %v\033[0m\n", err)
		return
	}

	var tasks []structs.Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Printf("\033[31mError parsing tasks: %v\033[0m\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Completed", "Created At"})

	for slno, task := range tasks {
		if task.Completed {
			status := "\033[32m✔\033[0m"
			createdAt := task.Created_at.Format("02 Jan 2006, 03:04 PM")
			table.Append([]string{
				strconv.Itoa(slno),
				task.Name,
				status,
				createdAt,
			})

			table.SetBorder(true)
			table.SetRowLine(true)
			table.SetAlignment(tablewriter.ALIGN_CENTER)

		}
	}
	table.Render()
}
func ListUndone() {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("\033[31mError reading tasks: %v\033[0m\n", err)
		return
	}

	var tasks []structs.Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		fmt.Printf("\033[31mError parsing tasks: %v\033[0m\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Completed", "Created At"})

	for slno, task := range tasks {

		if !task.Completed {
			status := "\033[31m✗\033[0m"

			createdAt := task.Created_at.Format("02 Jan 2006, 03:04 PM")
			table.Append([]string{
				strconv.Itoa(slno),
				task.Name,
				status,
				createdAt,
			})

			table.SetBorder(true)
			table.SetRowLine(true)
			table.SetAlignment(tablewriter.ALIGN_CENTER)

		}
	}
	table.Render()
}

func ResetAll() {
	fmt.Print("\033[33mAre you sure you want to reset all tasks? This action cannot be undone. (y/n): \033[0m")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "y" {
		fmt.Println("Resetting all tasks...")
		tasks := []structs.Task{}
		err := auxiliary.WriteFile(filePath, tasks)
		if err != nil {
			fmt.Println("\033[31mAction failed-Error.\033[0m")
		}
		fmt.Println("Action completed.")
	} else {
		fmt.Println("\033[31mAction cancelled.\033[0m")
	}
}
