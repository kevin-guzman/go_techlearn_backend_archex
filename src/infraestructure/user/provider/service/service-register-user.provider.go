package service

import (
	repositoryCompany "golang-gingonic-hex-architecture/src/domain/company/port/repository"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/domain/user/service"
)

func GetServiceRegisterUser(ru repository.RepositoryUser, rc repositoryCompany.RepositoryCompany) *service.ServiceRegisterUser {
	return service.NewServiceRegisterUser(ru, rc)
}
