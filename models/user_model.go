package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mohamedafify/go-backend/database"
	"github.com/mohamedafify/go-backend/database/oracle"
	"github.com/mohamedafify/go-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    []byte    `json:"password"`
	Token       string    `json:"token"`
	CreatedAt   time.Time
}

func CreateUser(req *SignupRequest) (*User, error) {
	// hash password
	encpw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating password when creating user, err: %v", err.Error())
		return nil, err
	}

	token, err := createToken()
	if err != nil {
		return nil, err
	}

	user := &User{
		Id:        uuid.New(),
		Email:     req.Email,
		Password:  encpw,
		Token:     token,
		CreatedAt: time.Now(),
	}

	oracleDb := oracle.OracleDatabase{}
	password, _ := utils.Getenv("DB_PASSWORD", "")
	connectionParams := map[string]any{
		"ip":       "localhost",
		"port":     1521,
		"service":  "XEPDB1",
		"user":     "DEMO",
		"password": password,
	}
	oracleDb.SetConnectionString(connectionParams)

	if err := createDatabaseUser(user, &oracleDb); err != nil {
		return nil, err
	}
	return user, nil
}

func createDatabaseUser(user *User, mydb db.Database) error {
	query := `
	INSERT INTO USERS(PK, EMAIL, PASSWORD, NAME, TOKEN)
	VALUES (:PK, :EMAIL, :PASSWORD, :NAME, :TOKEN)`

	values := []any{user.Id.String(), user.Email, string(user.Password), string(user.Name), user.Token}
	if err := db.Exec(mydb, query, values); err != nil {
		return err
	}
	return nil
}
