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
	handleEditUser      command.HandlerEditUser
	handleDeleteUser    command.HandlerDeleteUser
}

func NewControllerUser(hru command.HandlerRegisterUser, hlu query.HandlerListUsers, hlou command.HandlerLoginUser, heu command.HandlerEditUser, hdu command.HandlerDeleteUser) *ControllerUser {
	return &ControllerUser{
		handlerRegisterUser: hru,
		handlerListUsers:    hlu,
		handleLoginUser:     hlou,
		handleEditUser:      heu,
		handleDeleteUser:    hdu,
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

func (cu *ControllerUser) Update(command command.CommandEditUser) (string, error, int) {
	return cu.handleEditUser.Run(command)
}

func (cu *ControllerUser) Delete(command command.CommandDeleteUser) (string, error, int) {
	return cu.handleDeleteUser.Run(command)
}
