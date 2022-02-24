package service

import (
	"golang-gingonic-hex-architecture/src/domain/company/port/repository"
	"golang-gingonic-hex-architecture/src/domain/company/service"
)

func GetServiceRegisterCompany(rc repository.RepositoryCompany) *service.ServiceRegisterCompany {
	return service.NewServiceRegisterCompany(rc)
}
