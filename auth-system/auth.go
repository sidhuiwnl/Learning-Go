package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("super-secret-key")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string)error {
	return  bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateJWT(userId int) (string,error){
	claims := jwt.MapClaims{
		"user_id" : userId,
		"exp"  : time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	

	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userId int) (string,error){
	claims := jwt.MapClaims{
		"user_id" : userId,
		"exp"  : time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	

	return token.SignedString(jwtSecret)

}




