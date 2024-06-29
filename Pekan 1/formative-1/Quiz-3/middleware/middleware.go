package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, pwd, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak boleh kosong"})
			return
		}

		if uname == "admin" && pwd == "password" {
			c.Next()
			return
		}

		if uname == "editor" && pwd == "secret" {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak sesuai"})
	}
}
