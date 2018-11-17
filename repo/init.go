package repo

import (
	"github.com/jinzhu/gorm"
)

//repository instance
type RepoInstance struct {
	db *gorm.DB
}



func NewRepository(db *gorm.DB) *RepoInstance {
	return &RepoInstance{db}
}

