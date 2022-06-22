package middleware

import (
	"arquitetura-go/pkg/web"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set token environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "Token vazio"),
			)
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.DecodeError(http.StatusUnauthorized, "Token inválido"),
			)
			return
		}

		c.Next()
	}
}
