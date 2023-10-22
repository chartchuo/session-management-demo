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

type RefreshClaims struct {
	TokenID TokenID `json:"jti"`
	model.User
	jwt.RegisteredClaims
}

// DO NOT use in production
// Store data in databas or external cache instead of inmemory
var counterCache *cache.Cache
var revokedCache *cache.Cache

func init() {
	counterCache = cache.New(refreshExp, time.Hour)
	revokedCache = cache.New(refreshExp, time.Hour)
}

// New refresh claim.
// Issue at = Now()
func NewRefreshClaims(u *model.User) (rc *RefreshClaims) {
	rc = &RefreshClaims{User: *u, TokenID: *NewTokenID()}
	rc.IssuedAt = &jwt.NumericDate{Time: now().Add(refreshExp)}
	counterCache.Set(rc.TokenID.NUID, rc.TokenID.Counter, cache.DefaultExpiration)
	return rc.UpdateTime()
}

// Get refresh claims from context.
// Return error if not valid.
func NewRefreshClaimsFromContext(c *gin.Context, option ...jwt.ParserOption) (*RefreshClaims, error) {
	tokenString := common.ExtractJWT(c)
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.RefreshSecret), nil
	}, option...)

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
	now := now()
	rc.ExpiresAt = &jwt.NumericDate{Time: now.Add(refreshExp)}
	rc.NotBefore = &jwt.NumericDate{Time: now.Add(refreshNBF)}
	return rc
}

// Generate new token
func (rc *RefreshClaims) Rotate() (refreshTokenString string, accessTokenString string, err error) {
	now := now()

	// check expire
	if !now.Before(rc.ExpiresAt.Time) {
		return "", "", fmt.Errorf("token expired: %s", rc.ExpiresAt.Time.String())
	}
	// check not before
	if now.Before(rc.NotBefore.Time) {
		return "", "", fmt.Errorf("token not before: %s", rc.NotBefore.Time.String())
	}
	// check revoked counter
	_, found := revokedCache.Get(rc.TokenID.NUID)
	if found {
		return "", "", fmt.Errorf("invalid token counter: %s", rc.TokenID.String())
	}

	// check counter
	r, found := counterCache.Get(rc.TokenID.NUID)
	if found && r.(int) > rc.TokenID.Counter {
		//revok this rc
		rc.Revoke()
		return "", "", fmt.Errorf("invalid token counter: %s", rc.TokenID.String())
	}

	// rotate token
	rc.TokenID.Rotate()
	// add to counter cache
	counterCache.Set(rc.TokenID.NUID, rc.TokenID.Counter, cache.DefaultExpiration)

	rc.UpdateTime()

	return rc.JwtString()
}

func (rc *RefreshClaims) JwtString() (refreshTokenString string, accessTokenString string, err error) {

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	refreshTokenString, err = refreshToken.SignedString([]byte(common.RefreshSecret))
	if err != nil {
		return
	}

	ac := NewAccessClaims(&rc.User)
	ac.IssuedAt = &jwt.NumericDate{Time: rc.IssuedAt.Time}
	ac.ExpiresAt = &jwt.NumericDate{Time: rc.IssuedAt.Time.Add(accessExp)}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessTokenString, err = accessToken.SignedString([]byte(common.AccessSecret))
	if err != nil {
		return
	}
	return
}

func (rc *RefreshClaims) Revoke() {
	// add to revoked cache
	revokedCache.Set(rc.TokenID.NUID, rc.TokenID.Counter, cache.DefaultExpiration)
}
