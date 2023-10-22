package token

import (
	"backend/model"
	"testing"
)

func TestACExpire(t *testing.T) {

	ac := NewAccessClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	expire := ac.IsExpired()
	if expire {
		t.Errorf("err: expect not expire but expired")
	}
	mock.forward(accessExp)
	expire = ac.IsExpired()
	if !expire {
		t.Errorf("err: expect expire but not")
	}

}
