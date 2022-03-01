package dao

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/application/article/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/article/entity"

	"gorm.io/gorm"
)

type DaoArticlePostgreSql struct {
	daoArticle *gorm.DB
}

func NewDaoArticlePostgreSql(conn *gorm.DB) *DaoArticlePostgreSql {
	return &DaoArticlePostgreSql{
		daoArticle: conn.Model(&entity.Article{}),
	}
}

func (dap *DaoArticlePostgreSql) List() []*dto.ArticleDto {
	var articles []*dto.ArticleDto
	dap.daoArticle.Raw("SELECT * FROM articles").Scan(&articles)
	return articles
}

func (dap *DaoArticlePostgreSql) ListByFilters(filters dto.FilterArticlesDto) []*dto.ArticleDto {
	var articles []*dto.ArticleDto
	fmt.Println("Filt", filters, len(articles))
	statement := dap.daoArticle
	if filters.Limit != 0 {
		statement.Limit(filters.Limit)
	}
	if filters.Offset != 0 {
		statement.Offset(filters.Offset)
	}
	for _, v := range filters.Type {
		fmt.Println("iteration", v.String())
		statement.Or("type = ?", v.String())
	}
	statement.Find(&articles)
	for _, v := range articles {
		fmt.Println("ar", &v.Type)
	}
	return articles
}
