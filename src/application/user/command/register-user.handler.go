package command

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

type HandlerRegisterUser struct {
	serviceRegisterUser service.ServiceRegisterUser
}

func NewHandlerRegisterUser(sru *service.ServiceRegisterUser) *HandlerRegisterUser {
	return &HandlerRegisterUser{
		serviceRegisterUser: *sru,
	}
}

func (hru *HandlerRegisterUser) Run(commandRU CommandRegisterUser) interface{} {
	user, err := model.NewUser(commandRU.Name, commandRU.Password, commandRU.Role, commandRU.Email, commandRU.CompanyId)
	if err != nil {
		return err
	}
	return hru.serviceRegisterUser.Run(*user)
}
