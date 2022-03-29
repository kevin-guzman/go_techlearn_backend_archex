package query

import (
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	"golang-gingonic-hex-architecture/src/domain/publication/port/dao"
)

type HandlerGetOneById struct {
	daoPublications dao.DaoPublication
}

func NewHandlerGetOneById(daoA dao.DaoPublication) *HandlerGetOneById {
	return &HandlerGetOneById{
		daoPublications: daoA,
	}
}

func (hla *HandlerGetOneById) Run(id int) *dto.PublicationDto {
	return hla.daoPublications.GetOneById(id)
}
