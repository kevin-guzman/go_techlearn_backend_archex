package repository

import (
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/company/port/repository"
	classRepository "golang-gingonic-hex-architecture/src/infraestructure/company/adaptor/repository"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceRepository.RepositoryCompany

var GetRepositoryCompany = func(conn *gorm.DB) *interfaceRepository.RepositoryCompany {
	once.Do(func() {
		ru := classRepository.NewRepositoryCompanyPostgreSql(conn)
		iru := interfaceRepository.RepositoryCompany(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
