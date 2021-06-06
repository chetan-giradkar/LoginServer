package middlewares

import (
	service "LoginServer/ServiceLayer"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRequest(js *service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER = "Bearer "
		authHeader := c.GetHeader("Authorization")
		log.Println(authHeader)
		token := authHeader[len(BEARER):]
		valid, validateError := js.ValidateToken(token)
		if validateError != nil {
			log.Printf("Error validating token: %v\n", validateError.Error())
			c.AbortWithStatus(http.StatusInternalServerError)

			return
		}

		if !valid.Valid {
			log.Println("Token not valid")
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}
}
