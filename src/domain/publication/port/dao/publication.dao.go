package dao

import "golang-gingonic-hex-architecture/src/application/publication/query/dto"

type DaoPublication interface {
	List() []*dto.PublicationDto
	ListByFilters(filters dto.FilterPublicationsDto) []*dto.PublicationDto
	Search(text string) []*dto.PublicationDto
	GetOneById(id int) *dto.PublicationDto
}
