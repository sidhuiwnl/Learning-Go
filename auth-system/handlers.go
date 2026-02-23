package main

import (
	"net/http"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
)

func SignUp(w http.ResponseWriter, r *http.Request){

	var req struct {
		Email string   `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	hashedPassword, err := HashPassword(req.Password)
	
	if err != nil {
		http.Error(w, "Error hashing password", 500)
		return
	}

	result,err := DB.Exec(
		"INSERT INTO users(email,password) VALUES(?, ?)",
		req.Email,
		hashedPassword,
	)
	
	if err != nil {
		http.Error(w,err.Error(),400)
		return
	}

	id, _ := result.LastInsertId()

	accessToken, _ := GenerateJWT(int(id))


    refreshToken, _ := GenerateRefreshToken(int(id))


	json.NewEncoder(w).Encode(map[string]string{
		"refreshToken" : refreshToken,
		"accessToken" : accessToken,

	})


}


func Login(w http.ResponseWriter, r *http.Request){

	var req struct {
		Email  string  `json:"email"`
		Password string  `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	var user User

	err := DB.QueryRow(
		"SELECT * from users WHERE email = ?",
		req.Email,
	).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	if CheckPassword(req.Password, user.Password) != nil {
		http.Error(w, "Invalid credentials", 401)
		return
	}

	accessToken, _ := GenerateJWT(user.ID)
	refreshToken, _ := GenerateRefreshToken(user.ID)

	json.NewEncoder(w).Encode(map[string]string{
		"accessToken" : accessToken,
		"refreshToken" : refreshToken,
	})

}

func RefreshToken(w http.ResponseWriter, r *http.Request){

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	token,err := jwt.Parse(req.RefreshToken,func(token *jwt.Token)(interface {}, error){
		return jwtSecret,nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid refresh token", 401)
		return
	}
	
	claims := token.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	newAccessTokne, _ := GenerateJWT(userId)

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": newAccessTokne,
	})
}

func Profile(w http.ResponseWriter, r *http.Request){
	userId := r.Context().Value("user_id").(int)

	var user User

	DB.QueryRow(
		"SELECT id, email FROM users WHERE id = ?",
		userId,
	).Scan(&user.ID, &user.Email)

	json.NewEncoder(w).Encode(user)
}