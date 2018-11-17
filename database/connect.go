package database

import "github.com/jinzhu/gorm"

func Connect() (db *gorm.DB, err error) {
	db,err = gorm.Open("mysql", "root:db_root_password@tcp(db:3306)/penguin_bbs?charset=utf8&parseTime=True&loc=Local")
	return
}
