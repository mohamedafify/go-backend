package models

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func CreateUser(c *gin.Context, req *SignupRequest) (*User, error) {
	// hash password
	encpw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	token, err := createToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return nil, err
	}

	user := &User{
		Id:        uuid.New(),
		Email:     req.Email,
		Password:  encpw,
		Token:     token,
		CreatedAt: time.Now(),
	}
	if err := createDatabaseUser(c, user); err != nil {
		return nil, err
	}
	return user, nil
}

func createDatabaseUser(c *gin.Context, user *User) error {
	oracleDb := oracle.OracleDatabase{}
	password, _ := utils.Getenv("DB_PASSWORD", "")
	connectionParams := map[string]any{
		"ip":       "localhost",
		"port":     1521,
		"service":  "XEPDB1",
		"user":     "DEMO",
		"password": password,
	}
	oracleDb.ConnectionString(connectionParams)

	query := `
	INSERT INTO USERS(PK, EMAIL, PASSWORD, NAME, TOKEN)
	VALUES (:PK, :EMAIL, :PASSWORD, :NAME, :TOKEN)`

	values := []any{user.Id.String(), user.Email, string(user.Password), string(user.Name), user.Token}
	if err := db.Exec(c, &oracleDb, query, values); err == nil {
		return err
	}
	return nil
}
