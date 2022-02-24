package command

import (
	"golang-gingonic-hex-architecture/src/domain/company/model"
	"golang-gingonic-hex-architecture/src/domain/company/service"
	"net/http"
)

type HandlerRegisterCompany struct {
	serviceRegisterCompany service.ServiceRegisterCompany
}

func NewHandlerRegisterCompany(sru *service.ServiceRegisterCompany) *HandlerRegisterCompany {
	return &HandlerRegisterCompany{
		serviceRegisterCompany: *sru,
	}
}

func (hrc *HandlerRegisterCompany) Run(commandRC CommandRegisterCompany) (string, error, int) {
	company, err := model.NewCompany(commandRC.Name, commandRC.Owner, commandRC.Phone, commandRC.Email)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	message, err, status := hrc.serviceRegisterCompany.Run(*company)
	return message, err, status
}
