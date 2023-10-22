package token

import (
	"backend/model"
	"log"
	"testing"
	"time"
)

func TestRCRevoke(t *testing.T) {

	rc := NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshNBF)
	_, _, err := rc.Rotate()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	rc.Revoke()
	_, _, err = rc.Rotate()
	if err == nil {
		t.Errorf("expect error but return sucess")
	}

}
func TestRCExpire(t *testing.T) {

	var start = now()
	rc := NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshExp)
	mock.forward(-time.Second)
	_, _, err := rc.Rotate()
	var end = now()
	if err != nil {
		log.Printf("start: %v\n", start)
		log.Printf("end  : %v\n", end)
		t.Errorf("err: %v", err)
	}

	start = now()
	rc = NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshExp)
	_, _, err = rc.Rotate()
	end = now()
	if err == nil {
		log.Printf("start            : %v\n", start)
		log.Printf("end              : %v\n", end)
		log.Printf("rc.ExpiresAt.Time: %v\n", rc.ExpiresAt.Time)

		t.Errorf("expect error but return sucess")
	}
}

func TestRCNotBefore(t *testing.T) {

	var start = now()
	rc := NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshNBF)
	_, _, err := rc.Rotate()
	var end = now()
	if err != nil {
		log.Printf("start            : %v\n", start)
		log.Printf("end              : %v\n", end)
		log.Printf("rc.NotBefore.Time: %v\n", rc.NotBefore.Time)

		t.Errorf("err: %v", err)
	}
	if err != nil {
		log.Printf("start: %v\n", start)
		log.Printf("end: %v\n", end)
		t.Errorf("err: %v", err)
	}

	start = now()
	rc = NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshNBF)
	mock.forward(-time.Second)
	_, _, err = rc.Rotate()
	end = now()
	if err == nil {
		log.Printf("start            : %v\n", start)
		log.Printf("end              : %v\n", end)
		log.Printf("rc.NotBefore.Time: %v\n", rc.NotBefore.Time)

		t.Errorf("expect error but return sucess")
	}

}
func TestRCRotate(t *testing.T) {

	var start = now()
	rc := NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshExp - time.Second)
	_, _, err := rc.Rotate()
	var end = now()
	if err != nil {
		log.Printf("start: %v\n", start)
		log.Printf("end: %v\n", end)
		t.Errorf("err: %v", err)
	}

	start = now()
	rc = NewRefreshClaims(&model.User{UserID: "test", FirstName: "firstName", LastName: "lastName", Role: "user"})
	mock.forward(refreshExp - time.Second)
	_, _, err = rc.Rotate()
	end = now()
	if err != nil {
		log.Printf("start: %v\n", start)
		log.Printf("end: %v\n", end)
		t.Errorf("err: %v", err)
	}

}
