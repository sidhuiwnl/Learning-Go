package main

import (
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"context"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing token", 401)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", 401)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		ctx := context.WithValue(r.Context(), "user_id", userID)

		next(w, r.WithContext(ctx))
	}
}