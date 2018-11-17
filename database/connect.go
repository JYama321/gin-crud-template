package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (db *gorm.DB, err error) {
	db,err = gorm.Open("mysql", "root:db_root_password@tcp(db:3306)/gin_template?charset=utf8&parseTime=True&loc=Local")
	return
}
