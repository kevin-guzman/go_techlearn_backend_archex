package service

import (
	"golang-gingonic-hex-architecture/src/domain/article/port/repository"
	"golang-gingonic-hex-architecture/src/domain/article/service"
)

func GetServiceCreateArticle(ra repository.RepositoryArticle) *service.ServiceCreateArticle {
	return service.NewServiceCreateArticle(ra)
}
