package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Segments   []Segment `gorm:"many2many:user_tags;"`
}

type Segment struct {
	gorm.Model
	Name string `json:"name"`
}
