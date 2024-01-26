package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mohamedafify/go-backend/utils"
)

type LoginRequest struct {
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

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginRequest := &LoginRequest{}
		if err := utils.BindBody(c, loginRequest); err != nil {
			return
		}
		token, err := createToken()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqApiKey := c.GetHeader("x-api-key")
		if len(reqApiKey) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "x-api-key header is required"})
			return
		}
		if utils.API_KEY == reqApiKey {
			reqJwtToken := c.GetHeader("authorization")
			if len(reqJwtToken) == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "authorization header is required"})
				return
			}
			token, err := jwt.Parse(reqJwtToken, func(token *jwt.Token) (interface{}, error) {
				return utils.JWT_SECRET, nil
			})
			if err != nil {
				log.Print(err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token"})
				return
			}
			if !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token"})
				return
			}
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid api key"})
			return
		}
	}
}
