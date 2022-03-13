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

func (hdu *HandlerDeleteUser) Run(commandDU CommandDeleteUser) (string, error, int) {
	user := model.User{Email: commandDU.Email, Password: commandDU.Password}
	return hdu.serviceDeleteUser.Run(commandDU.UserId, user)
}
