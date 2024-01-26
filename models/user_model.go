package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	Token       string    `json:"token"`
	CreatedAt   time.Time
}

func CreateUser() *User {
	user := &User{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
	}

	return user
}
