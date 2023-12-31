package common

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// DO NOT use in production
// Store secure key in vault instead
const RefreshSecret = "refresh secret"
const AccessSecret = "access secret"

// Extract JWT from gin context.
// Header: Bearer XXX
// return "" if not found
func ExtractJWT(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")

	return strings.TrimPrefix(authHeader, "Bearer ")
}
