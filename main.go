package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamedafify/go-backend/handlers"
	"github.com/mohamedafify/go-backend/utils"
)

func setupServer() *gin.Engine {
	server := gin.Default()
	utils.SetupValidations(server)
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	server.SetTrustedProxies(nil)
	return server
}

func main() {
	logFile := utils.SetupLoggers()
	defer (*logFile).Close()
	utils.SetupEnv()
	server := setupServer()
	handlers.Setup(server)
	port, _ := utils.Getenv("PORT", "8080")
	ip := fmt.Sprintf("0.0.0.0:%v", port)
	log.Fatal(server.Run(ip))
}
