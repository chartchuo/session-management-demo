package service

import (
	"backend/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshTokenHandler(c *gin.Context) {

	rc, err := token.NewRefreshClaimsFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		return
	}

	refreshTokenString, accessTokenString, err := rc.Rotate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refresh_token": refreshTokenString,
		"access_token":  accessTokenString,
	})

}
