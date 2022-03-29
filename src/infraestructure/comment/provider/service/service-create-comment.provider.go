package service

import (
	"golang-gingonic-hex-architecture/src/domain/comment/port/repository"
	"golang-gingonic-hex-architecture/src/domain/comment/service"
	publicationRepository "golang-gingonic-hex-architecture/src/domain/publication/port/repository"
)

func GetServiceCreateComment(rc repository.RepositoryComment, rp publicationRepository.RepositoryPublication) *service.ServiceCreateComment {
	return service.NewServiceCreateComment(rc, rp)
}
