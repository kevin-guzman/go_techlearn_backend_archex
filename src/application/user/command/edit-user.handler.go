package command

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

type HandlerEditUser struct {
	serviceEditUser service.ServiceEditUser
}

func NewHandlerEditUser(seu *service.ServiceEditUser) *HandlerEditUser {
	return &HandlerEditUser{
		serviceEditUser: *seu,
	}
}

func (heu *HandlerEditUser) Run(commandEU CommandEditUser) interface{} {
	user := model.User{Name: commandEU.Name, Email: commandEU.Email}
	return heu.serviceEditUser.Run(commandEU.UserId, user)
}
