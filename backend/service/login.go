package service

import (
	"backend/model"
	"backend/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid login values",
		})
		return
	}
	userID := loginVals.Username
	password := loginVals.Password
	var u *model.User

	// DONT use in production
	// store data in database instead
	if userID == "admin" && password == "admin" {
		u = &model.User{
			Role:      "admin",
			UserID:    userID,
			FirstName: "Admin",
			LastName:  "Admin",
		}
	}
	if userID == "test" && password == "test" {
		u = &model.User{
			Role:      "user",
			UserID:    userID,
			FirstName: "Test",
			LastName:  "Test",
		}
	}
	if u == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalud username or password",
		})
		return
	}
	rc := token.NewRefreshClaims(u)
	ac := token.NewAccessClaims(u)
	refreshTokenString, err := rc.JwtString()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	accessTokenString, err := ac.JwtString()
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
