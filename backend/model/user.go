package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Role      string `json:"role"`
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Get User data from context.
// Router must under auth middleware.
// Return error if not found.
func UserFromContext(c *gin.Context) (*User, error) {
	userValue, exist := c.Get("user")
	if !exist {
		return nil, fmt.Errorf("user context not exist")
	}
	return userValue.(*User), nil
}
