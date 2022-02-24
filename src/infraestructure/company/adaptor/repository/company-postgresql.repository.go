package repository

import (
	"golang-gingonic-hex-architecture/src/domain/company/model"
	"golang-gingonic-hex-architecture/src/infraestructure/company/entity"

	"gorm.io/gorm"
)

type RepositoryCompanyPostgreSql struct {
	companyRepository *gorm.DB
}

func NewRepositoryCompanyPostgreSql(conn *gorm.DB) *RepositoryCompanyPostgreSql {
	return &RepositoryCompanyPostgreSql{
		companyRepository: conn.Model(&entity.Company{}),
	}
}

func (rcp *RepositoryCompanyPostgreSql) Save(company model.Company) error {
	entity := entity.Company{Name: company.Name, Owner: company.Owner, Email: company.Email, Phone: company.Phone}
	result := rcp.companyRepository.Create(&entity)
	return result.Error
}

func (rcp *RepositoryCompanyPostgreSql) ExistCompanyByName(name string) (bool, error) {
	var company model.Company
	var count int64 = 0
	result := rcp.companyRepository.Find(&company, "name = ?", name).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (rcp *RepositoryCompanyPostgreSql) ExistCompanyById(id int) (bool, error) {
	var company model.Company
	var count int64 = 0
	result := rcp.companyRepository.Find(&company, id).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
