package main

import (
	"backend/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ac, err := token.NewAccessClaimsFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
		} else {
			c.Set("user", &ac.User)
		}

	}
}
