package dao

import (
	interfaceDao "golang-gingonic-hex-architecture/src/domain/article/port/dao"
	classDao "golang-gingonic-hex-architecture/src/infraestructure/article/adaptor/dao"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceDao.DaoArticle

var GetDaoArticle = func(conn *gorm.DB) *interfaceDao.DaoArticle {
	once.Do(func() {
		ru := classDao.NewDaoArticlePostgreSql(conn)
		iru := interfaceDao.DaoArticle(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
