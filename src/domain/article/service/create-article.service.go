package service

import (
	"golang-gingonic-hex-architecture/src/domain/article/model"
	"golang-gingonic-hex-architecture/src/domain/article/port/repository"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"net/http"
)

var (
	errTrace       string = "This error has ocurred in create-article.service.go"
	internalError  string = "Internal server error"
	successMessage string = "Article has succesfully created!"
)

type ServiceCreateArticle struct {
	articleRepository repository.RepositoryArticle
}

func NewServiceCreateArticle(ArticleR repository.RepositoryArticle) *ServiceCreateArticle {
	return &ServiceCreateArticle{
		articleRepository: ArticleR,
	}
}

func (sca *ServiceCreateArticle) Run(article model.Article) (string, error, int) {
	err := sca.articleRepository.Save(article)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}
