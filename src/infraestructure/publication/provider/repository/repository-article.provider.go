package repository

import (
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/publication/port/repository"
	classRepository "golang-gingonic-hex-architecture/src/infraestructure/publication/adaptor/repository"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceRepository.RepositoryPublication

var GetRepositoryPublication = func(conn *gorm.DB) *interfaceRepository.RepositoryPublication {
	once.Do(func() {
		ru := classRepository.NewRepositoryPublicationPostgreSql(conn)
		iru := interfaceRepository.RepositoryPublication(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
