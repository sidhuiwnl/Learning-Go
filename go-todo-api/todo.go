package main

type Todo struct {
	Id  	int     `json:"id"`
	Name    string  `json:"name"`
	Status  bool    `json:"status"`
}