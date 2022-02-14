package entity

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type Password string

type User struct {
	gorm.Model
	Id            int `gorm:"primaryKey"`
	Name          string
	Password      string
	Creation_date time.Time
	Role          string
	Email         string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return fmt.Errorf("Error crypting the user password %v", err)
	}
	u.Password = string(bytes)
	return nil
}
