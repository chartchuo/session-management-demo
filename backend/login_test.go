package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	reader := strings.NewReader("{\"username\":\"test\",\"password\":\"test\"}")

	req, _ := http.NewRequest("POST", "/login", reader)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("error:%v", w.Body.String())
	}

	w = httptest.NewRecorder()
	reader = strings.NewReader("{\"username\":\"admin\",\"password\":\"admin\"}")

	req, _ = http.NewRequest("POST", "/login", reader)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("error:%v", w.Body.String())
	}

	w = httptest.NewRecorder()
	reader = strings.NewReader("{\"username\":\"user\",\"password\":\"incorrect password\"}")

	req, _ = http.NewRequest("POST", "/login", reader)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	if w.Code == 200 {
		t.Errorf("error:%v", "expect error but sucess")
	}
}
