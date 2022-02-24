package entity

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/infraestructure/company/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id            int `gorm:"primaryKey"`
	Name          string
	Password      string
	Creation_date time.Time
	Role          string
	Email         string
	CompanyId     int
	Company       entity.Company `gorm:"foreignKey:CompanyId"`
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
