package controller

import (
	"golang-gingonic-hex-architecture/src/application/company/command"
	"golang-gingonic-hex-architecture/src/application/company/query"
	"golang-gingonic-hex-architecture/src/application/company/query/dto"
)

type ControllerCompany struct {
	handlerRegisterCompany command.HandlerRegisterCompany
	handlerListCompanies   query.HandlerListCompanies
}

func NewControllerCompany(hrc command.HandlerRegisterCompany, hlc query.HandlerListCompanies) *ControllerCompany {
	return &ControllerCompany{
		handlerRegisterCompany: hrc,
		handlerListCompanies:   hlc,
	}
}

func (cu *ControllerCompany) Create(command command.CommandRegisterCompany) (string, error, int) {
	return cu.handlerRegisterCompany.Run(command)
}

func (cu *ControllerCompany) List() []*dto.CompanyDto {
	return cu.handlerListCompanies.Run()
}
