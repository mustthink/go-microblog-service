package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	AuthorID uint64 `gorm:"index"`
	TopicID  uint64 `gorm:"index"`
	Title    string `gorm:"size:255"`
	Content  string `gorm:"type:text"`
}
