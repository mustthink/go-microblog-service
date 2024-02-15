package models

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	Title   string `gorm:"size:255"`
	Content string `gorm:"type:text"`
}
