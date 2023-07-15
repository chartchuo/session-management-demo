package main

import (
	"backend/model"
	"testing"
)

func TestAuthorizeUser(t *testing.T) {
	u := &model.User{UserID: "test", Role: "user"}
	//positive
	_, err := authorize(u, "GET", "/user/test/hello")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	//negative
	_, err = authorize(u, "GET", "/user/test2/hello")
	if err == nil {
		t.Errorf("err:")
	}
	_, err = authorize(u, "GET", "/admin/hello")
	if err == nil {
		t.Errorf("err:")
	}
	_, err = authorize(u, "GET", "/admin/hello")
	if err == nil {
		t.Errorf("err:")
	}
	_, err = authorize(u, "GET", "/admin")
	if err == nil {
		t.Errorf("err:")
	}
	_, err = authorize(u, "GET", "/other")
	if err == nil {
		t.Errorf("err:")
	}

}

func TestAuthorizeAdmin(t *testing.T) {
	u := &model.User{UserID: "adminUser", Role: "admin"}
	//positive
	_, err := authorize(u, "GET", "/user/test/hello")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	_, err = authorize(u, "GET", "/user/test2/hello")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	_, err = authorize(u, "GET", "/admin/hello")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	//negative
	_, err = authorize(u, "GET", "/admin")
	if err == nil {
		t.Errorf("err: ")
	}
	_, err = authorize(u, "GET", "/other")
	if err == nil {
		t.Errorf("err: ")
	}

}
