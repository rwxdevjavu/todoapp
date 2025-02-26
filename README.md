# todoapp

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
