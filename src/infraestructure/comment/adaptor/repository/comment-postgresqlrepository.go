package repository

import (
	"golang-gingonic-hex-architecture/src/domain/comment/model"
	"golang-gingonic-hex-architecture/src/infraestructure/comment/entity"

	"gorm.io/gorm"
)

type RepositoryComentPostgreSql struct {
	publicationRepository *gorm.DB
}

func NewRepositoryComentPostgreSql(conn *gorm.DB) *RepositoryComentPostgreSql {
	return &RepositoryComentPostgreSql{
		publicationRepository: conn.Model(&entity.Comment{}),
	}
}

func (rcp *RepositoryComentPostgreSql) Save(comment model.Comment) error {
	entity := entity.Comment{
		CommentTitle:   comment.Title,
		CommentContent: comment.Content,
		UserId:         comment.UserId,
		PublicationId:  comment.PublicationId,
		WrittenAt:      comment.WrittenAt,
	}
	result := rcp.publicationRepository.Create(&entity)
	return result.Error
}
