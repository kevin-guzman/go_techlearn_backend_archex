package command

import (
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"golang-gingonic-hex-architecture/src/domain/publication/service"
)

type HandlerCreatePublication struct {
	serviceCreatePublication service.ServiceCreatePublication
}

func NewHandlerCreatePublication(sca *service.ServiceCreatePublication) *HandlerCreatePublication {
	return &HandlerCreatePublication{
		serviceCreatePublication: *sca,
	}
}

func (hca *HandlerCreatePublication) Run(command CommandCreatePublication) interface{} {
	article, err := model.NewPublication(command.Title, command.Description, command.Content, command.WiterUserId, command.Type, command.Categories)
	if err != nil {
		return err
	}
	return hca.serviceCreatePublication.Run(*article)
}
