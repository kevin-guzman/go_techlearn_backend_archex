package dao

import "golang-gingonic-hex-architecture/src/application/company/query/dto"

type DaoCompany interface {
	List() []*dto.CompanyDto
}
