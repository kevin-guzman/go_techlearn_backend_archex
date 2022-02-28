package dao

import "golang-gingonic-hex-architecture/src/application/article/query/dto"

type DaoArticle interface {
	List() []*dto.ArticleDto
	ListByFilters(filters dto.FilterArticlesDto) []*dto.ArticleDto
}
