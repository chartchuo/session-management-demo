package token

import (
	"fmt"
	"time"

	"backend/common"
	"backend/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/patrickmn/go-cache"
)

const refreshExp = time.Minute * 2
const refreshNBF = time.Second * 30
const accessExp = time.Second * 60

type RefreshClaims struct {
	TokenID TokenID `json:"jti"`
	model.User
	jwt.RegisteredClaims
}

var counterCache *cache.Cache
var invalidCache *cache.Cache

func init() {
	counterCache = cache.New(refreshExp, time.Hour)
	invalidCache = cache.New(refreshExp, time.Hour)
}

// New refresh claim.
// Issue at = Now()
func NewRefreshClaims(u *model.User) (rc *RefreshClaims) {
	rc = &RefreshClaims{User: *u, TokenID: *NewTokenID()}
	return rc.UpdateTime()
}

// Get refresh claims from context.
// Return error if not valid.
func NewRefreshClaimsFromContext(c *gin.Context) (*RefreshClaims, error) {
	tokenString := common.ExtractJWT(c)
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.RefreshSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %v", err)
	}
	rc, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token: %v", "invlid claims")
	}

	return rc, nil
}

// Update iat,exp,nbf from current time
func (rc *RefreshClaims) UpdateTime() *RefreshClaims {
	iat := time.Now()
	rc.IssuedAt = &jwt.NumericDate{Time: iat}
	rc.ExpiresAt = &jwt.NumericDate{Time: iat.Add(refreshExp)}
	rc.NotBefore = &jwt.NumericDate{Time: iat.Add(refreshNBF)}
	return rc
}

// Generate new token
func (rc *RefreshClaims) Rotate() (refreshTokenString string, accessTokenString string, err error) {

	// check invalid counter
	_, found := invalidCache.Get(rc.TokenID.NUID)
	if found {
		return "", "", fmt.Errorf("invalid token counter: %s", rc.TokenID.String())
	}

	// check counter
	r, found := counterCache.Get(rc.TokenID.NUID)
	if found && r.(int) > rc.TokenID.Counter {
		// add to invalid cache
		invalidCache.Set(rc.TokenID.NUID, rc.TokenID.Counter, cache.DefaultExpiration)
		return "", "", fmt.Errorf("invalid token counter: %s", rc.TokenID.String())
	}

	// rotate token
	rc.TokenID.Rotate()
	// add to counter cache
	counterCache.Set(rc.TokenID.NUID, rc.TokenID.Counter, cache.DefaultExpiration)

	rc.UpdateTime()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	refreshTokenString, err = refreshToken.SignedString([]byte(common.RefreshSecret))
	if err != nil {
		return
	}

	ac := NewAccessClaims(&rc.User)
	iat := rc.IssuedAt
	ac.IssuedAt = iat
	ac.ExpiresAt = &jwt.NumericDate{Time: iat.Add(accessExp)}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessTokenString, err = accessToken.SignedString([]byte(common.AccessSecret))
	if err != nil {
		return
	}
	return
}
