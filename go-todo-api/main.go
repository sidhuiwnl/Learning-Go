package main

import (
	"log"
	"net/http"
	"fmt"
)


func main(){
	InitDB()

	http.HandleFunc("/todos",func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			GetTodos(w,r)

		}else if r.Method == http.MethodPost {
			createTodo(w,r)
		}

	})


	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, world!")
	})

	http.HandleFunc("/todo",func (w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")

			fmt.Fprintf(w, "ID is: %s",id)
		}
	})

	log.Println("Server runing on :8080")

	http.ListenAndServe(":8080",nil)
}