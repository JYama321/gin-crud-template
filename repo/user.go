package repo

import (
	"fmt"
	"github.com/JYama321/gin-crud-template/entity"
)

/*
 get user response code
	0: successfully created
	1: already exist
	-1: error has occurred
*/
type CreateUserResponseCode int8

/*
get user response code
	0: successfully
	-1: error has occurred
*/

type GetUserResponseCode int8



func (r *RepoInstance) getByID(id int) entity.User {
	var user entity.User
	r.db.First(&user, id)
	return user
}


func (r *RepoInstance) GetUsers() ([]entity.User, GetUserResponseCode ,error) {
	var user []entity.User
	if err := r.db.Limit(10).Find(&user).Error; err != nil {
		fmt.Println(err)
		return user, 1, err
	} else {
		return user, -1, nil
	}
}

func (r *RepoInstance) CreateUser(
	userName string,
	email string,
	passwordDigest string,
) (entity.User, CreateUserResponseCode,error) {
	var user = entity.User{
		UserName: userName,
		Email: email,
		PasswordDigest:passwordDigest,
	}
	if r.db.NewRecord(&user) {
		// 新規登録
		result := r.db.Create(&user)
		result.Scan(&user)
		return user, 0, nil
	} else {
		// すでにユーザが存在
		return user, 1, nil
	}
}
