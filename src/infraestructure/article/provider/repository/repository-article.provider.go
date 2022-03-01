package repository

import (
	interfaceRepository "golang-gingonic-hex-architecture/src/domain/article/port/repository"
	classRepository "golang-gingonic-hex-architecture/src/infraestructure/article/adaptor/repository"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceRepository.RepositoryArticle

var GetRepositoryArticle = func(conn *gorm.DB) *interfaceRepository.RepositoryArticle {
	once.Do(func() {
		ru := classRepository.NewRepositoryArticlePostgreSql(conn)
		iru := interfaceRepository.RepositoryArticle(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
