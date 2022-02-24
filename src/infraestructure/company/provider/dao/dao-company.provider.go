package dao

import (
	interfaceDao "golang-gingonic-hex-architecture/src/domain/company/port/dao"
	classDao "golang-gingonic-hex-architecture/src/infraestructure/company/adaptor/dao"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var instance *interfaceDao.DaoCompany

var GetDaoCompany = func(conn *gorm.DB) *interfaceDao.DaoCompany {
	once.Do(func() {
		ru := classDao.NewDaoCompanyPostgreSql(conn)
		iru := interfaceDao.DaoCompany(ru)
		_, _ = iru, ru
		instance = &iru
	})
	return instance
}
