package utils

import (
	"log"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var API_KEY string
var JWT_SECRET []byte

func SetupEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to get enviroment variables with err: %v", err.Error())
	}

	ginMode, _ := Getenv("GIN_MODE", gin.DebugMode)
	gin.SetMode(ginMode)

	// validate existence of an API_KEY for validation
	apiKey, ok := Getenv("API_KEY", "")
	if !ok {
		log.Fatal("Must insert API_KEY in your enviroment variables")
	}
	API_KEY = apiKey

	jwtSecret, ok := Getenv("JWT_SECRET", "")
	if !ok {
		log.Fatal("Must insert JWT_SECRET in your enviroment variables")
	}
	JWT_SECRET = []byte(jwtSecret)
}

func Getenv(key string, defaultValue string) (string, bool) {
	if result, ok := syscall.Getenv(key); ok && len(result) != 0 {
		return result, true
	}
	log.Printf("Failed to get %v from your enviroment", key)
	return defaultValue, false
}
