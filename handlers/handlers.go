package handlers

import (
	"github.com/gin-gonic/gin"
)

func Setup(server *gin.Engine) {
	server.POST("/login", Login())
	authorized := server.Group("/auth")
	authorized.Use(Auth())
	authorized.GET("/user", handleUser)
}
