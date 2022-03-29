package dao

import (
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"

	"gorm.io/gorm"
)

type DaoUserPostgreSql struct {
	daoUser *gorm.DB
}

func NewDaoUserPostgreSql(conn *gorm.DB) *DaoUserPostgreSql {
	return &DaoUserPostgreSql{
		daoUser: conn.Model(&entity.User{}),
	}
}

func (dup *DaoUserPostgreSql) List() []*dto.UserDto {
	var users []*dto.UserDto
	dup.daoUser.Raw("SELECT * FROM USERS u WHERE u.deleted_at IS NULL").Scan(&users)
	return users
}
