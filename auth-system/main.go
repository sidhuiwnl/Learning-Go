package main

import (
	"net/http"
	"fmt"
)

func main (){
	
	InitDB()

	http.HandleFunc("/signup",SignUp)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/refreshToken", RefreshToken)
	http.HandleFunc("/profile", AuthMiddleware(Profile))

	fmt.Println("The server is running on port 8080")
	http.ListenAndServe(":8080",nil)

}



