package repository

import "golang-gingonic-hex-architecture/src/domain/company/model"

type RepositoryCompany interface {
	ExistCompanyById(id int) (bool, error)
	ExistCompanyByName(name string) (bool, error)
	Save(company model.Company) error
}
