package query

import (
	"golang-gingonic-hex-architecture/src/application/company/query/dto"
	"golang-gingonic-hex-architecture/src/domain/company/port/dao"
)

type HandlerListCompanies struct {
	daoCompany dao.DaoCompany
}

func NewHandlerListCompanies(daoU dao.DaoCompany) *HandlerListCompanies {
	return &HandlerListCompanies{
		daoCompany: daoU,
	}
}

func (hlc *HandlerListCompanies) Run() []*dto.CompanyDto {
	return hlc.daoCompany.List()
}
