package model

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"time"
)

const (
	MIN_PASSWORD_LENGTH int    = 8
	ERR_LENGTH          string = "La contraseña es muy corta"
)

type User struct {
	Name          string
	Password      string
	Creation_date time.Time
	Id            int
	Role          string
	Email         string
	CompanyId     int
}

func NewUser(name, password, role, email string, companyId int) (*User, *errors.ErrorCore) {
	if len(password) < MIN_PASSWORD_LENGTH {
		err := fmt.Errorf(ERR_LENGTH)
		return nil, errors.NewErrorInvalidLength(err, err.Error())
	}
	return &User{
		Name:      name,
		Password:  password,
		Role:      role,
		Email:     email,
		CompanyId: companyId,
	}, nil
}
