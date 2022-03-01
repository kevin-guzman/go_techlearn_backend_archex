package repository

import (
	"golang-gingonic-hex-architecture/src/domain/article/model"
	"golang-gingonic-hex-architecture/src/infraestructure/article/entity"

	"gorm.io/gorm"
)

type RepositoryArticlePostgreSql struct {
	articleRepository *gorm.DB
}

func NewRepositoryArticlePostgreSql(conn *gorm.DB) *RepositoryArticlePostgreSql {
	return &RepositoryArticlePostgreSql{
		articleRepository: conn.Model(&entity.Article{}),
	}
}

func (rcp *RepositoryArticlePostgreSql) Save(article model.Article) error {
	entity := entity.Article{
		Title:       article.Title,
		Description: article.Description,
		Content:     article.Content,
		WiterUserId: article.WiterUserId,
		WrittenAt:   article.WrittenAt,
		Type:        article.Type,
	}
	result := rcp.articleRepository.Create(&entity)
	return result.Error
}
