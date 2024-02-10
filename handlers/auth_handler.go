package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mohamedafify/go-backend/models"
	"github.com/mohamedafify/go-backend/utils"
)

type authClaims struct {
	jwt.StandardClaims
	userId uuid.UUID
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		signupRequest := &models.SignupRequest{}
		if err := utils.BindBody(c, signupRequest); err != nil {
			return
		}
		user, err := models.CreateUser(c, signupRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": user.Token})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 		loginRequest := &models.LoginRequest{}
		// 		if err := utils.BindBody(c, loginRequest); err != nil {
		// 			return
		// 		}
		// 		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func setContextKeys(c *gin.Context, data *authClaims) {
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
			myClaims := &authClaims{}
			token, err := jwt.ParseWithClaims(reqJwtToken, myClaims, func(token *jwt.Token) (interface{}, error) {
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
			setContextKeys(c, myClaims)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid api key"})
			return
		}
	}
}
