package repository

import (
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"golang-gingonic-hex-architecture/src/infraestructure/publication/entity"

	"gorm.io/gorm"
)

type RepositoryPublicationPostgreSql struct {
	publicationRepository *gorm.DB
}

func NewRepositoryPublicationPostgreSql(conn *gorm.DB) *RepositoryPublicationPostgreSql {
	return &RepositoryPublicationPostgreSql{
		publicationRepository: conn.Model(&entity.Publication{}),
	}
}

func (rcp *RepositoryPublicationPostgreSql) Save(publication model.Publication) error {
	entity := entity.Publication{
		Title:       publication.Title,
		Description: publication.Description,
		Content:     publication.Content,
		WiterUserId: publication.WiterUserId,
		WrittenAt:   publication.WrittenAt,
		Type:        publication.Type,
	}
	result := rcp.publicationRepository.Create(&entity)
	return result.Error
}
