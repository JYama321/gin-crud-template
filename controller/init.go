package controller

import repository2 "github.com/JYama321/gin-crud-template/repository"

var repository repository2.RepoInstance
var repo *repository2.RepoInstance

func init() () {
	repository = repository2.NewRepository()
	repo = &repository
}
