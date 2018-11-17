package repository

import (
	"github.com/jinzhu/gorm"
)

//repository instance
type RepoInstance struct {
	db *gorm.DB
}

/*
 get user response code
	0: successfully created
	1: already exist
	-1: error has occurred
 */
type createUserResponseCode int8

/*
get user response code
	0: successfully
	-1: error has occurred
*/

type getUserResponseCode int8


func NewRepository() RepoInstance {
	return RepoInstance{}
}

