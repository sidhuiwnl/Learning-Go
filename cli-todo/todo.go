package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func main() {

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask(os.Args[2:])
	case "list":
		todos, err := loadTodos()

		if err != nil {
			fmt.Println("Couldn't able to load Todos")
		}

		if len(todos) <= 0 {
			fmt.Println("No Todos Left, Great!")
		} else {
			for _, t := range todos {

				status := "❌"

				if t.Status {
					status = "✅"
				}

				fmt.Printf("%d. %s %s\n", t.Id, t.Name, status)
			}
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo complete <id>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		completeTodo(id)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
	}

}

func printHelp() {
	fmt.Println(`

	Todo App - Simple task manager

	todo add <task>      Add a new task
	todo list            Show all tasks
	todo complete <id>   Mark task as done
	todo delete <id>     Remove a task
	todo help            Show this help

	`)
}

func completeTodo(todoId int) {

	todos, _ := loadTodos()

	for i := range todos {
		if todos[i].Id == todoId {
			todos[i].Status = true
			saveTodos(todos)
			fmt.Println("Task completed")
			return
		}
	}

	fmt.Println("Task not found")
}

func loadTodos() ([]Todo, error) {

	var todos []Todo

	data, err := os.ReadFile("todo.txt")

	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}

		return nil, err
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return nil, err
	}

	return todos, nil

}

func saveTodos(todos []Todo) {

	data, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		fmt.Println("Error converting JSON:", err)
		return
	}

	err = os.WriteFile("todo.txt", data, 0644)

	if err != nil {
		fmt.Println("Error writing file:", err)
	}

}

func addTask(args []string) {

	if len(args) < 1 {
		fmt.Println("Error: Task description required")
		fmt.Println("Usage: todo add <task description>")
		return
	}

	task := strings.Join(args, " ")

	todos, err := loadTodos()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newId := 1

	if len(todos) > 0 {
		newId = todos[len(todos)-1].Id + 1
	}

	newTodo := Todo{
		Id:     newId,
		Name:   task,
		Status: false,
	}

	todos = append(todos, newTodo)

	saveTodos(todos)

}
