package service

import (
	"net/http"

	"backend/model"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	u, err := model.UserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"user_id":    u.UserID,
		"first_name": u.FirstName,
		"text":       "Hello World.",
	})
}
