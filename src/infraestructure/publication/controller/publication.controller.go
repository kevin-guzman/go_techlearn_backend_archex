package controller

import (
	"golang-gingonic-hex-architecture/src/application/publication/command"
	"golang-gingonic-hex-architecture/src/application/publication/query"
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
)

type ControllerPublication struct {
	handlerCreatePublication      command.HandlerCreatePublication
	handlerListPublication        query.HandlerListPublication
	handlerListFiltredPublication query.HandlerListFiltredPublication
}

func NewControllerPublication(hcp command.HandlerCreatePublication, hlp query.HandlerListPublication, hlfp query.HandlerListFiltredPublication) *ControllerPublication {
	return &ControllerPublication{
		handlerCreatePublication:      hcp,
		handlerListPublication:        hlp,
		handlerListFiltredPublication: hlfp,
	}
}

func (cp *ControllerPublication) Create(command command.CommandCreatePublication) (string, error, int) {
	return cp.handlerCreatePublication.Run(command)
}

func (cp *ControllerPublication) List() []*dto.PublicationDto {
	return cp.handlerListPublication.Run()
}

func (cp *ControllerPublication) ListFiltred(filters dto.FilterPublicationsDto) []*dto.PublicationDto {
	return cp.handlerListFiltredPublication.Run(filters)
}
