package dao

import (
	interfaceDao "golang-gingonic-hex-architecture/src/domain/publication/port/dao"
	classDao "golang-gingonic-hex-architecture/src/infraestructure/publication/adaptor/dao"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceDao.DaoPublication

var GetDaoPublication = func(conn *gorm.DB) *interfaceDao.DaoPublication {
	once.Do(func() {
		ru := classDao.NewDaoPublicationPostgreSql(conn)
		iru := interfaceDao.DaoPublication(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
