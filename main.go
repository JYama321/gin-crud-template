package main

import (
	"fmt"
	"github.com/JYama321/gin-crud-template/controller"
	"github.com/JYama321/gin-crud-template/database"
	"github.com/JYama321/gin-crud-template/entity"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func gormConnect() *gorm.DB {
	fmt.Println("start connecting...")
	db, err := database.Connect()
	db.AutoMigrate(&entity.User{})
	if err != nil {
		fmt.Printf("database connection error %s", err.Error())
		panic("database connection failed.")
	}
	return db
}

func main() {
	router := gin.Default()

	router.GET("/users", controller.GetUsers)
	router.POST("/users/new", controller.NewUser)

	db := gormConnect()
	defer db.Close()


	router.Run(":3000")
}