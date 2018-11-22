package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http/httptest"
	"testing"
)

func TestNewUserSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/users/new",NewUser)

	var body newUserRequest
	body.Email = "yamaji@gmail.com"
	body.Username = "Jiro Yamamoto"
	body.Password = "JiroPassword"

	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/users/new", bytes.NewReader(b))
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestGetUsersSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/users",GetUsers)
	req := httptest.NewRequest("GET", "/users",nil)
	r.ServeHTTP(w,req)

	if w.Code != 200 {
		t.Fail()
	}
}