package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedafify/go-backend/models"
)

func handleUser(c *gin.Context) {
	user := models.CreateUser()
	c.JSON(http.StatusOK, user)
}
