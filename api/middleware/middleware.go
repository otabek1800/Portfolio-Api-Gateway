package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/otabek1800/Portfolio-Api-Gateway/api/handler"
)

func AuthMiddleware(h *handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CORSMiddleware(h *handler.Handler) gin.HandlerFunc {	
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func ErrorMiddleware(h *handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
