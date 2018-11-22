package repo

import (
	"github.com/JYama321/gin-crud-template/database"
	"github.com/jinzhu/gorm"
)

func connect() *gorm.DB {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	return db
}

