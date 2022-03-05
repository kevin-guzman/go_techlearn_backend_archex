package service

import (
	"golang-gingonic-hex-architecture/src/domain/publication/port/repository"
	"golang-gingonic-hex-architecture/src/domain/publication/service"
)

func GetServiceCreatePublication(rp repository.RepositoryPublication) *service.ServiceCreatePublication {
	return service.NewServiceCreatePublication(rp)
}
