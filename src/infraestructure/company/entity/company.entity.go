package entity

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Id    int `gorm:"primaryKey"`
	Name  string
	Owner string
	Phone string
	Email string
}
