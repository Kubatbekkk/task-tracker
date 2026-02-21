package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Kubatbekkk/task-tracker/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Description required")
			return
		}

		t, err := task.AddTask(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", t.ID)

	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		err := task.UpdateTask(id, os.Args[3])
		handleErr(err)

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		err := task.DeleteTask(id)
		handleErr(err)

	case "mark-in-progress":
		id, _ := strconv.Atoi(os.Args[2])
		err := task.ChangeStatus(id, task.StatusInProgress)
		handleErr(err)

	case "mark-done":
		id, _ := strconv.Atoi(os.Args[2])
		err := task.ChangeStatus(id, task.StatusDone)
		handleErr(err)

	case "list":
		if len(os.Args) == 2 {
			list("")
			return
		}
		list(task.Status(os.Args[2]))

	default:
		fmt.Println("Unknown Command")
	}
}

func list(status task.Status) {
	tasks, err := task.ListTasks(status)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, t := range tasks {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Success")
}
