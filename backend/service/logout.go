package service

import (
	"backend/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LogoutHandler(c *gin.Context) {

	rc, err := token.NewRefreshClaimsFromContext(c, jwt.WithoutClaimsValidation())
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
