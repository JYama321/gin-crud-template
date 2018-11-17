package controller

import (
	"github.com/JYama321/gin-crud-template/repo"
	"github.com/jinzhu/gorm"
)

var rep *repo.RepoInstance

func NewInstance(db *gorm.DB) () {
	rep = repo.NewRepository(db)
}
