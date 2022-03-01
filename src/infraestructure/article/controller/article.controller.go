package controller

import (
	"golang-gingonic-hex-architecture/src/application/article/command"
	"golang-gingonic-hex-architecture/src/application/article/query"
	"golang-gingonic-hex-architecture/src/application/article/query/dto"
)

type ControllerArticle struct {
	handlerCreateArticle       command.HandlerCreateArticle
	handlerListArticles        query.HandlerListArticles
	handlerListFiltredArticles query.HandlerListFiltredArticles
}

func NewControllerArticle(hca command.HandlerCreateArticle, hla query.HandlerListArticles, hlfa query.HandlerListFiltredArticles) *ControllerArticle {
	return &ControllerArticle{
		handlerCreateArticle:       hca,
		handlerListArticles:        hla,
		handlerListFiltredArticles: hlfa,
	}
}

func (cu *ControllerArticle) Create(command command.CommandCreateArticle) (string, error, int) {
	return cu.handlerCreateArticle.Run(command)
}

func (cu *ControllerArticle) List() []*dto.ArticleDto {
	return cu.handlerListArticles.Run()
}

func (cu *ControllerArticle) ListFiltred(filters dto.FilterArticlesDto) []*dto.ArticleDto {
	return cu.handlerListFiltredArticles.Run(filters)
}
