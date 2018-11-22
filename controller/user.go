package controller

import (
	"github.com/JYama321/gin-crud-template/entity"
	"github.com/JYama321/gin-crud-template/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newUserRequest struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type newUserResponse struct {
	User entity.User `json:"user"`
	Code repo.CreateUserResponseCode `json:"resultCode"`
	Error error `json:"error"`
}

type getUserResponse struct {
	Users []entity.User `json:"users"`
	Code repo.GetUserResponseCode `json:"resultCode"`
	Error error `json:"error"`
}


/*
	/users request
*/
func GetUsers(c *gin.Context) {
	var res getUserResponse

	rep := repo.NewRepository()
	users,code,err := rep.GetUsers()
	res.Users = users
	res.Code = code
	res.Error = err
	if res.Error == nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, res)
	}
}


/*
	/users/new request
*/

func NewUser(c *gin.Context) {
	var json newUserRequest

	rep := repo.NewRepository()
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var res newUserResponse
	user,code,err := rep.CreateUser(json.Username,json.Email,json.Password)
	res.User = user
	res.Code = code
	res.Error = err
	c.JSON(http.StatusOK, res)
}

