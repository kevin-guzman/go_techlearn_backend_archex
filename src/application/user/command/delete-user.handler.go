package command

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

type HandlerDeleteUser struct {
	serviceDeleteUser service.ServiceDeleteUser
}

func NewHandlerDeleteUser(sdu *service.ServiceDeleteUser) *HandlerDeleteUser {
	return &HandlerDeleteUser{
		serviceDeleteUser: *sdu,
	}
}

func (hdu *HandlerDeleteUser) Run(id int, commandDU CommandDeleteUser) interface{} {
	user := model.User{Email: commandDU.Email, Password: commandDU.Password}
	return hdu.serviceDeleteUser.Run(id, user)
}
