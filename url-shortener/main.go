package main

import (
	"log"
	"net/http"
)

func main(){
	InitDB()

	http.HandleFunc("/shorten",CreateShortUrl)

	http.HandleFunc("/",RedirectUrl)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080",nil)

}