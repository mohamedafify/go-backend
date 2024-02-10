package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/mohamedafify/go-backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{})
	tokenString, err := token.SignedString(utils.JWT_SECRET)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
