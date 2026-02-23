package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func GetTodos(w http.ResponseWriter, r *http.Request ){

	rows, err := DB.Query("SELECT * from `todos`")

	fmt.Println("The value directly from query",rows)

	if err != nil {
		http.Error(w,err.Error(),500)
	}

	defer rows.Close()

	todos := []Todo{}

	for rows.Next(){
		var t Todo

		rows.Scan(&t.Id, &t.Name, &t.Status)

		todos = append(todos, t)

	}

	json.NewEncoder(w).Encode(todos)
}


func createTodo(w http.ResponseWriter, r *http.Request ){

	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result, err := DB.Exec(`INSERT into todos(name,status) values (?, ?)`, todo.Name,todo.Status)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, _ := result.LastInsertId()

	todo.Id = int(id)

	json.NewEncoder(w).Encode(todo)

}