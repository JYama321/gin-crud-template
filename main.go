package main

import (
	"fmt"
	"github.com/JYama321/gin-crud-template/database"
	"github.com/JYama321/gin-crud-template/entity"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func gormConnect() *gorm.DB {
	db, err := database.Connect()
	db.AutoMigrate(&entity.User{})

	if err != nil {
		fmt.Println("database connection error %s", err)
		panic("database connection failed.")
	}
	return db
}

func main() {
	router := gin.Default()

	router.GET("/users")
	router.POST("/users/new")
}