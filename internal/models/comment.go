package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	PostID   uint64 `gorm:"index"`
	AuthorID uint64
	Content  string `gorm:"type:text"`
}
