package repository

import (
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"golang-gingonic-hex-architecture/src/infraestructure/publication/entity"
	"strconv"

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
		Categories:  publication.Categories,
	}
	result := rcp.publicationRepository.Create(&entity)
	return result.Error
}

func (rcp *RepositoryPublicationPostgreSql) ExistById(id int) (bool, error) {
	// var user model.Publication
	var count int64 = 0

	parsedId := strconv.Itoa(id)
	rawSql := "SELECT p.id FROM publications p WHERE p.id = " + parsedId
	result := rcp.publicationRepository.Raw(rawSql).Count(&count)

	// result := rcp.publicationRepository.Where("id = ?", id).Find(&user).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}
