package middleware

import "github.com/gin-gonic/gin"

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		return
	}
}
