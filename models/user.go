package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}
