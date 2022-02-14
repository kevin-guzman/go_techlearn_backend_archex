package controller

import (
	"golang-gingonic-hex-architecture/src/application/user/command"
	"golang-gingonic-hex-architecture/src/application/user/query"
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
)

type ControllerUser struct {
	handlerRegisterUser command.HandlerRegisterUser
	handlerListUsers    query.HandlerListUsers
	handleLoginUser     command.HandlerLoginUser
}

func NewControllerUser(hru command.HandlerRegisterUser, hlu query.HandlerListUsers, hlou command.HandlerLoginUser) *ControllerUser {
	return &ControllerUser{
		handlerRegisterUser: hru,
		handlerListUsers:    hlu,
		handleLoginUser:     hlou,
	}
}

func (cu *ControllerUser) Create(command command.CommandRegisterUser) (string, error, int) {
	return cu.handlerRegisterUser.Run(command)
}

func (cu *ControllerUser) List() []*dto.UserDto {
	return cu.handlerListUsers.Run()
}

func (cu *ControllerUser) Login(command command.CommandLoginUser) (string, error, int) {
	return cu.handleLoginUser.Run(command)
}
