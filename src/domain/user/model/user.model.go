package model

import (
	"fmt"
	"time"
)

const MIN_PASSWORD_LENGTH = 6

type User struct {
	Name          string
	Password      string
	Creation_date time.Time
	Id            int
	Role          string
	Email         string
	CompanyId     int
}

func NewUser(name, password, role, email string, companyId int) (*User, error) {
	if len(password) < MIN_PASSWORD_LENGTH {
		return nil, fmt.Errorf("The leng of the password is incorrect")
	}
	return &User{
		Name:      name,
		Password:  password,
		Role:      role,
		Email:     email,
		CompanyId: companyId,
	}, nil
}
