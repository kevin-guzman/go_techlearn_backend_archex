package command

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
	"net/http"
)

type HandlerRegisterUser struct {
	serviceRegisterUser service.ServiceRegisterUser
}

func NewHandlerRegisterUser(sru *service.ServiceRegisterUser) *HandlerRegisterUser {
	return &HandlerRegisterUser{
		serviceRegisterUser: *sru,
	}
}

func (hru *HandlerRegisterUser) Run(commandRU CommandRegisterUser) (string, error, int) {
	user, err := model.NewUser(commandRU.Name, commandRU.Password, commandRU.Role, commandRU.Email, commandRU.CompanyId)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	message, err, status := hru.serviceRegisterUser.Run(*user)
	return message, err, status
}
