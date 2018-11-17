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
	fmt.Println("start connecting...")
	db, err := database.Connect()
	fmt.Printf("error %s",err)
	db.AutoMigrate(&entity.User{})


	if err != nil {
		fmt.Println("database connection error %s", err)
		panic("database connection failed.")
	}
	defer db.Close()

	controller.NewInstance(db)
	router.GET("/users", controller.GetUsers)
	router.POST("/users/new", controller.NewUser)

	router.Run(":3000")
}