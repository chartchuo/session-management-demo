package token

import (
	"backend/common"
	"backend/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	model.User
	jwt.RegisteredClaims
}

func NewAccessClaims(u *model.User) (ac *AccessClaims) {
	ac = &AccessClaims{User: *u}
	return
}

// Verify access token.
// Return error if not found or invalid.
func NewAccessClaimsFromContext(c *gin.Context) (*AccessClaims, error) {
	tokenString := common.ExtractJWT(c)
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.AccessSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid access token: %v", err)
	}
	ac, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid access token: %v", "invlid claims")
	}

	return ac, nil
}
