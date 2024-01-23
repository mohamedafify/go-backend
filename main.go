package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamedafify/go-backend/utils"
)

func setupRouter() *gin.Engine {
	server := gin.Default()
	return server
}

func main() {
	logFile := utils.SetupLoggers()
	defer (*logFile).Close()
	utils.SetupEnv()
	server := setupRouter()
	server.SetTrustedProxies(nil)
	log.Fatal(server.Run(":8080"))
}
