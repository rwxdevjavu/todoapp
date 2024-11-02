package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/rwxdevjavu/todoapp/crud"
)

func helpPrompt() {
	fmt.Print(`
Usage:
  todoapp [command] [flags]

Commands:
  add       Add a new task to your todo list.
  list      List all tasks in your todo list.
  done      Mark a task as completed.
  delete    Delete a task from your todo list.
  help      Display this help message.

Flags:
  -t, --title   The title of the task (used with 'add' and 'done' commands).
  -d, --desc    Description of the task (used with 'add' command).
  -i, --id      The ID of the task (used with 'done' and 'delete' commands).
  -h, --help    Display this help message.

Examples:
  todo add -t "Buy groceries" -d "Milk, Eggs, Bread"
      Adds a new task with the title "Buy groceries" and description "Milk, Eggs, Bread."

  todo list
      Lists all tasks with their ID, title, status, and description.

  todo done -i 2
      Marks the task with ID 2 as completed.

  todo delete -i 3
      Deletes the task with ID 3 from your todo list.

  todo help
      Displays this help message.

Notes:
  - Task IDs are assigned automatically when you add a new task.
  - Completed tasks will be indicated in the task list.
  - Use quotes around multi-word titles or descriptions.

Happy task managing!`)
}

func main() {
	var todos crud.Todos
	todos.LoadJSON("todo.json")
	if len(os.Args) <= 1 {
		helpPrompt()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		taskTitle := addCmd.String("t", "", "todoapp : title of the task to add. See 'todoapp help'")
		taskDescription := addCmd.String("d", "", "todoapp : description of the task to add (optional). See 'todoapp help'")
		addCmd.Parse(os.Args[2:])
		if *taskTitle == "" {
			fmt.Println("Task Title cannot be empty. See 'todoapp help'")
			os.Exit(1)
		}
		todos.AddTask(*taskTitle, *taskDescription)
		todos.SaveJSON("todo.json")
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listCmd.Parse(os.Args[2:])
		tbl := table.New("id", "title", "description", "completed")
		//
		for i := range todos {
			var status string
			if todos[i].Status == true {
				status = "🟢"
			} else {
				status = "🔴"
			}
			//status :=
			tbl.AddRow(i, todos[i].Title, todos[i].Description, status)
		}
		tbl.Print()
	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		taskId := doneCmd.Int("i", -1, "todoapp : id of the task to mark as done. See 'todoapp help'")
		doneCmd.Parse(os.Args[2:])
		if *taskId < 0 {
			fmt.Println("todoapp :id cannot be empty")
			os.Exit(1)
		}
		if err := todos.MarkCompleted(*taskId); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		todos.SaveJSON("todo.json")
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		taskId := deleteCmd.Int("i", -1, "todoapp : id of the task to delete. See 'todoapp help'")
		deleteCmd.Parse(os.Args[2:])
		if *taskId < 0 {
			fmt.Println("todoapp :id cannot be empty")
			os.Exit(1)
		}
		if err := todos.RemoveTask(*taskId); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		todos.SaveJSON("todo.json")
	case "help":
		helpPrompt()
	default:
		fmt.Println("todoapp: " + os.Args[1] + " is not a todoapp command. See 'todoapp help'")
	}
}
