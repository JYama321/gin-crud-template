package controller

import (
	"github.com/JYama321/gin-crud-template/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newUserRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type getUserRequest struct {
	UserName string `json:"username"`
	Email string `json:"email"`

}

type getUserResponse struct {
	User entity.User `json:"user"`
	Code int `json:"resultCode"`
	Error error `json:"error"`
}


type newUserResponse struct {
	UserName string `json:"user"`
	Code int `json:"resultCode"`
	Error error `json:"error"`
}


func GetUsers(c *gin.Context) {
	var res getUserResponse
	res = repo.GetUsers()
	if res.Error := nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, res)
	}
}


func NewUser(c *gin.Context) {
	var json newUserRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var res newUserResponse
	res = repo.createUser(&json)
	c.JSON(http.StatusOK, res)
}

