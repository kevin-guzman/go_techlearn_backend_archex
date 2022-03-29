package repository

import (
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/comment/port/repository"
	classRepository "golang-gingonic-hex-architecture/src/infraestructure/comment/adaptor/repository"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceRepository.RepositoryComment

var GetRepositoryComment = func(conn *gorm.DB) *interfaceRepository.RepositoryComment {
	once.Do(func() {
		ru := classRepository.NewRepositoryComentPostgreSql(conn)
		iru := interfaceRepository.RepositoryComment(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
