package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func CreateShortUrl(w http.ResponseWriter, r *http.Request){

	var req struct {
		URL string  `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.URL == "" {
		http.Error(w,"Invalid request body",400)
		return
	}

	shortCode := generateShortCode(6)

	_, err = DB.Exec(
		"INSERT INTO url(original_url, short_code) VALUES (?, ?)",
		req.URL,
		shortCode,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]string{
		"short_url": "http://localhost:8080/" + shortCode,
	}

	// response = {
	//  short_url = "http://localhost:8080/dadadad"    
	// }

	json.NewEncoder(w).Encode(response)
}


func RedirectUrl(w http.ResponseWriter, r *http.Request){

	shortCode := r.URL.Path[1:]

	var originalURL string

	err := DB.QueryRow(
		"SELECT original_url FROM url WHERE short_code = ?",
		shortCode,
	).Scan(&originalURL)

	if err != nil {
		http.NotFound(w,r)
		return
	}

	fmt.Println("the code",originalURL)

	http.Redirect(w,r,originalURL,http.StatusFound)
	
	
}