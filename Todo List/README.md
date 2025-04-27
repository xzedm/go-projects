Todo List CLI
A simple command-line Todo application written in Go that allows users to add, list, and mark tasks as done. Tasks are stored in a JSON file for persistence.
Features

Add tasks with a description
List all tasks with their status (done or not)
Mark tasks as done by their number
Persist tasks to todos.json
Load tasks from todos.json on startup

Prerequisites

Go 1.16 or higher

Installation

Clone or download the project.
Navigate to the project directory.
Run the program:go run todo.go



Usage
Run the program and use the following commands:

add <description>: Add a new task (e.g., add Buy groceries).
list: Display all tasks with their numbers and status.
done <task_number>: Mark a task as done (e.g., done 1).
exit: Quit the application.

Example session:
> list
No tasks in the list.
> add Buy groceries
Task added: Buy groceries
> add Call Alice
Task added: Call Alice
> list
1. [ ] Buy groceries
2. [ ] Call Alice
> done 1
Task 1 marked as done.
> list
1. [✓] Buy groceries
2. [ ] Call Alice
> exit
Goodbye!

File Structure

todo.go: Main source code
todos.json: Stores tasks (created automatically)
README.md: This file

Notes

Tasks are saved to todos.json after each add or done command.
Invalid inputs (e.g., non-numeric task numbers) are handled with error messages.
The code follows Go best practices, including proper error handling and formatting with gofmt.

