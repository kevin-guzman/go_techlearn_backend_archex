package query

import (
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	"golang-gingonic-hex-architecture/src/domain/publication/port/dao"
)

type HandlerListFiltredPublication struct {
	daoPublications dao.DaoPublication
}

func NewHandlerListFiltredPublication(daoA dao.DaoPublication) *HandlerListFiltredPublication {
	return &HandlerListFiltredPublication{
		daoPublications: daoA,
	}
}

func (hla *HandlerListFiltredPublication) Run(filters dto.FilterPublicationsDto) []*dto.PublicationDto {
	return hla.daoPublications.ListByFilters(filters)
}
