package utils

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to get enviromentVariables")
	}

	if ginMode := os.Getenv("GIN_MODE"); strings.Trim(ginMode, " ") == "" {
		log.Print("Failed to get GIN_MODE in your enviroment")
		gin.SetMode(gin.DebugMode)
	}
}
