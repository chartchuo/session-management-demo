package service

import (
	"backend/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {

	rc, err := token.NewRefreshClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		return
	}

	rc.Revoke()

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})

}
