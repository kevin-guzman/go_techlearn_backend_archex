package controller

import (
	"golang-gingonic-hex-architecture/src/application/publication/command"
	"golang-gingonic-hex-architecture/src/application/publication/query"
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
)

type ControllerPublication struct {
	handlerCreatePublication        command.HandlerCreatePublication
	handlerListPublication          query.HandlerListPublication
	handlerListFiltredPublication   query.HandlerListFiltredPublication
	handlerListSearchedPublications query.HandlerListSearchedPublications
	handlerGetOneById               query.HandlerGetOneById
}

func NewControllerPublication(
	hcp command.HandlerCreatePublication,
	hlp query.HandlerListPublication,
	hlfp query.HandlerListFiltredPublication,
	hlsp query.HandlerListSearchedPublications,
	hgobid query.HandlerGetOneById,
) *ControllerPublication {
	return &ControllerPublication{
		handlerCreatePublication:        hcp,
		handlerListPublication:          hlp,
		handlerListFiltredPublication:   hlfp,
		handlerListSearchedPublications: hlsp,
		handlerGetOneById:               hgobid,
	}
}

func (cp *ControllerPublication) Create(command command.CommandCreatePublication) interface{} {
	return cp.handlerCreatePublication.Run(command)
}

func (cp *ControllerPublication) List() []*dto.PublicationDto {
	return cp.handlerListPublication.Run()
}

func (cp *ControllerPublication) ListFiltred(filters dto.FilterPublicationsDto) []*dto.PublicationDto {
	return cp.handlerListFiltredPublication.Run(filters)
}

func (cp *ControllerPublication) Search(text string) []*dto.PublicationDto {
	return cp.handlerListSearchedPublications.Run(text)
}

func (cp *ControllerPublication) GetOneById(id int) *dto.PublicationDto {
	return cp.handlerGetOneById.Run(id)
}
