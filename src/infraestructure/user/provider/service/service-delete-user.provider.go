package service

import (
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

func GetServiceDeleteUser(ru repository.RepositoryUser) *service.ServiceDeleteUser {
	return service.NewServiceDeleteUser(ru)
}
