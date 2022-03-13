package service

import (
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

func GetServiceEditUser(ru repository.RepositoryUser) *service.ServiceEditUser {
	return service.NewServiceEditUser(ru)
}
