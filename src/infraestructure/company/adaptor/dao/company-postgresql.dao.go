package dao

import (
	"golang-gingonic-hex-architecture/src/application/company/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/company/entity"

	"gorm.io/gorm"
)

type DaoCompanyPostgreSql struct {
	daoCompany *gorm.DB
}

func NewDaoCompanyPostgreSql(conn *gorm.DB) *DaoCompanyPostgreSql {
	return &DaoCompanyPostgreSql{
		daoCompany: conn.Model(&entity.Company{}),
	}
}

func (dcp *DaoCompanyPostgreSql) List() []*dto.CompanyDto {
	var companies []*dto.CompanyDto
	dcp.daoCompany.Raw("SELECT u.name, u.id FROM companies u").Scan(&companies)
	return companies
}
