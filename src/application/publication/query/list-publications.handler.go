package query

import (
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	"golang-gingonic-hex-architecture/src/domain/publication/port/dao"
)

type HandlerListPublication struct {
	daoPublications dao.DaoPublication
}

func NewHandlerListPublication(daoA dao.DaoPublication) *HandlerListPublication {
	return &HandlerListPublication{
		daoPublications: daoA,
	}
}

func (hla *HandlerListPublication) Run() []*dto.PublicationDto {
	return hla.daoPublications.List()
}
