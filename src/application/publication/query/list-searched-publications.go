package query

import (
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	"golang-gingonic-hex-architecture/src/domain/publication/port/dao"
)

type HandlerListSearchedPublications struct {
	daoPublications dao.DaoPublication
}

func NewHandlerLSearchedistPublications(daoA dao.DaoPublication) *HandlerListSearchedPublications {
	return &HandlerListSearchedPublications{
		daoPublications: daoA,
	}
}

func (hla *HandlerListSearchedPublications) Run(text string) []*dto.PublicationDto {
	return hla.daoPublications.Search(text)
}
