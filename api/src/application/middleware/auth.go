package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	jwtConfig "app/application/config/jwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthenticationRequired verifica se o token jwt informado é valido
func AuthenticationRequired(c *gin.Context) {
	if ok, err := strconv.ParseBool(jwtConfig.Disabled); ok && err == nil {
		c.Next()
		return
	}

	tHeader, has := c.Request.Header["Authorization"]
	if !has {
		log.Println("[AuthenticationRequired] Token not found")
		c.JSON(http.StatusUnauthorized, []gin.H{{"message": "user needs to be signed in to access this service"}})
		c.Abort()
		return
	}
	tokenString := tHeader[0]
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	tokenString = strings.Replace(tokenString, "bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	claims, ok := token.Claims.(*jwt.StandardClaims)

	if err != nil || !token.Valid || !ok || claims.Valid() != nil || claims.Id == "" {
		c.JSON(http.StatusUnauthorized, []gin.H{{"message": "user needs to be signed in to access this service"}})
		c.Abort()
		return
	}

	c.Next()

}
