package command

import (
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

type HandlerLoginUser struct {
	serviceLoginUser service.ServiceLoginUser
}

func NewHandlerLoginUser(slu *service.ServiceLoginUser) *HandlerLoginUser {
	return &HandlerLoginUser{
		serviceLoginUser: *slu,
	}
}

func (hlu *HandlerLoginUser) Run(commandLU CommandLoginUser) interface{} {
	return hlu.serviceLoginUser.Run(commandLU.Email, commandLU.Password)
}
