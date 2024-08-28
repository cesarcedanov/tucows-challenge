package server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

// This should NOT be pushed but for the Challenge purpose I did push it
var jwtSecretKey = []byte("eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTcyNDgwODY1MywiaWF0IjoxNzI0ODA4NjUzfQ.DPkUmBg7zmnDMeDRmai4-ORiiNF2D6G2jqBm2zwF6qI")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// Employee needs to check credentials hourly
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	})
	return token.SignedString(jwtSecretKey)
}

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"errorMsg": "No Authorization header"})
			c.Abort()
			return
		}
		tokenClaims := &Claims{}
		token, err := jwt.ParseWithClaims(authHeader, tokenClaims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"errorMsg": "Invalid Authorization header (token)"})
			c.Abort()
			return
		}

		c.Set("username", tokenClaims.Username)
		c.Next()
	}
}

func Login(c *gin.Context) {
	var userCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&userCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.ToLower(userCredentials.Username) != strings.ToLower(userCredentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
	}
	token, err := generateToken(userCredentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
