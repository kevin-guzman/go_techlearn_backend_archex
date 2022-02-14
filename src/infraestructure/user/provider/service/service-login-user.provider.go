package service

import (
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

func GetServiceLoginUser(ru repository.RepositoryUser) *service.ServiceLoginUser {
	return service.NewServiceLoginUser(ru)
}
