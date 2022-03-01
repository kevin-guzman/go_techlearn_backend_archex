package query

import (
	"golang-gingonic-hex-architecture/src/application/article/query/dto"
	"golang-gingonic-hex-architecture/src/domain/article/port/dao"
)

type HandlerListArticles struct {
	daoArticles dao.DaoArticle
}

func NewHandlerListArticles(daoA dao.DaoArticle) *HandlerListArticles {
	return &HandlerListArticles{
		daoArticles: daoA,
	}
}

func (hla *HandlerListArticles) Run() []*dto.ArticleDto {
	return hla.daoArticles.List()
}
