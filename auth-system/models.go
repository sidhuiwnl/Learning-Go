package main

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // "-" prevents sending password in json response
}