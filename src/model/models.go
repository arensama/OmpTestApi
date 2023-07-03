package model

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"-" `
	User   User   `json:"-" `
}
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname" `
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique_index;not null" `
	Blogs    []Blog `json:"blogs" gorm:"foreignKey:UserID"`
}
