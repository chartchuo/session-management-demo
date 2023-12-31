package main

import (
	"backend/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ac, err := token.ExtractAccessClaims(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		if _, err := authorize(&ac.User, c.Request.Method, c.Request.URL.Path); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "unauthorized to access: " + err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user", &ac.User)

	}
}
