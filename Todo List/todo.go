package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Todo represents a single task
type Todo struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// TodoList manages the collection of todos
type TodoList struct {
	Tasks []Todo `json:"tasks"`
}

// saveToFile writes the todo list to a JSON file
func (tl *TodoList) saveToFile(filename string) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}

// loadFromFile reads the todo list from a JSON file
func loadFromFile(filename string) (*TodoList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &TodoList{Tasks: []Todo{}}, nil // Return empty list if file doesn't exist
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var tl TodoList
	if err := json.Unmarshal(data, &tl); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return &tl, nil
}

// addTask adds a new task to the list
func (tl *TodoList) addTask(description string) {
	tl.Tasks = append(tl.Tasks, Todo{Description: description, Done: false})
}

// listTasks prints all tasks with their indices
func (tl *TodoList) listTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}
	for i, task := range tl.Tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

// markDone marks a task as done by its 1-based index
func (tl *TodoList) markDone(index int) error {
	if index < 1 || index > len(tl.Tasks) {
		return fmt.Errorf("invalid task number: %d", index)
	}
	tl.Tasks[index-1].Done = true
	return nil
}

func main() {
	const filename = "todos.json"
	todoList, err := loadFromFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading todos: %v\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Todo List CLI")
	fmt.Println("Commands: add <description>, list, done <task_number>, exit")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command := args[0]

		switch command {
		case "exit":
			fmt.Println("Goodbye!")
			return
		case "add":
			if len(args) < 2 {
				fmt.Println("Usage: add <description>")
				continue
			}
			description := strings.Join(args[1:], " ")
			todoList.addTask(description)
			if err := todoList.saveToFile(filename); err != nil {
				fmt.Println("Error saving todos:", err)
			} else {
				fmt.Println("Task added:", description)
			}
		case "list":
			todoList.listTasks()
		case "done":
			if len(args) != 2 {
				fmt.Println("Usage: done <task_number>")
				continue
			}
			index, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid task number:", args[1])
				continue
			}
			if err := todoList.markDone(index); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if err := todoList.saveToFile(filename); err != nil {
				fmt.Println("Error saving todos:", err)
			} else {
				fmt.Println("Task", index, "marked as done.")
			}
		default:
			fmt.Println("Unknown command. Use: add, list, done, exit")
		}
	}
}
