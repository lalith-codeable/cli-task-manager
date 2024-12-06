package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/lalith-codeable/cli-task-manager/auxiliary"
	"github.com/lalith-codeable/cli-task-manager/tasks"
)

func main() {
	auxiliary.Checks()
	if len(os.Args) < 2 {
		auxiliary.Scream()
		return
	}

	add := flag.NewFlagSet("add", flag.ExitOnError)
	list := flag.NewFlagSet("list", flag.ExitOnError)
	done := flag.NewFlagSet("done", flag.ExitOnError)
	undone := flag.NewFlagSet("undone", flag.ExitOnError)
	delete := flag.NewFlagSet("delete", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])
		if len(add.Args()) < 1 {
			fmt.Println("Required: A unique task name.")
			fmt.Println(`Usage: task add "<task name>"`)
			return
		}
		tasks.Add(add.Arg(0))
	case "list":
		list.Parse(os.Args[2:])
		if len(list.Args()) == 1 {
			verb := list.Arg(0)
			if verb == "done" {
				tasks.ListDone()
				return
			} else if verb == "undone" {
				tasks.ListUndone()
				return
			} else {
				fmt.Printf("Invalid list command verb :'%v'.", verb)
				fmt.Println("Usage: task list done || task list undone")
				return
			}
		}
		tasks.List()

	case "done":
		done.Parse(os.Args[2:])
		if len(done.Args()) > 0 {
			id, err := strconv.ParseUint(done.Arg(0), 10, 8)
			if err != nil {
				fmt.Println("Invalid command usage, pass a valid id.")
				fmt.Println("Usage: task done <task id>")
			}
			tasks.Toggle(uint8(id))
		} else {
			fmt.Println("Invalid command usage, pass a valid id.")
			fmt.Println("Usage: task done <task id>")
		}
	case "undone":
		undone.Parse(os.Args[2:])
		if len(undone.Args()) > 0 {
			id, err := strconv.ParseUint(undone.Arg(0), 10, 8)
			if err != nil {
				fmt.Println("Invalid command usage, pass a valid id.")
				fmt.Println("Usage: task undone <task id>")
			}
			tasks.Toggle(uint8(id))
		} else {
			fmt.Println("Invalid command usage, pass a valid id.")
			fmt.Println("Usage: task undone <task id>")
		}
	case "delete":
		delete.Parse(os.Args[2:])
		if len(delete.Args()) > 0 {
			var ids []uint8
			for _, id := range delete.Args() {
				ParedId, err := strconv.ParseUint(id, 10, 8)
				if err != nil {
					fmt.Printf("Invalid task id: %v ...", id)
				}
				ids = append(ids, uint8(ParedId))
			}

			tasks.Delete(ids)
		} else {
			fmt.Println("Invalid command usage, pass a valid id.")
			fmt.Println("Usage: task delete <task id>")
		}
	case "reset-all":
		tasks.ResetAll()

	default:
		auxiliary.Scream()
	}
}
