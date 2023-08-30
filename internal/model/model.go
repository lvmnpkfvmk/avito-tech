package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID string
	Tags   []Tag `gorm:"many2many:user_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
}
