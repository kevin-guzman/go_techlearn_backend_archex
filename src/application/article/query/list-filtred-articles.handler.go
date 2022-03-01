package query

import (
	"golang-gingonic-hex-architecture/src/application/article/query/dto"
	"golang-gingonic-hex-architecture/src/domain/article/port/dao"
)

type HandlerListFiltredArticles struct {
	daoArticles dao.DaoArticle
}

func NewHandlerListFiltredArticles(daoA dao.DaoArticle) *HandlerListFiltredArticles {
	return &HandlerListFiltredArticles{
		daoArticles: daoA,
	}
}

func (hla *HandlerListFiltredArticles) Run(filters dto.FilterArticlesDto) []*dto.ArticleDto {
	return hla.daoArticles.ListByFilters(filters)
}
