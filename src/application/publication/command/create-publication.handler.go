package command

import (
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"golang-gingonic-hex-architecture/src/domain/publication/service"
	"net/http"
)

type HandlerCreatePublication struct {
	serviceCreatePublication service.ServiceCreatePublication
}

func NewHandlerCreatePublication(sca *service.ServiceCreatePublication) *HandlerCreatePublication {
	return &HandlerCreatePublication{
		serviceCreatePublication: *sca,
	}
}

func (hca *HandlerCreatePublication) Run(command CommandCreatePublication) (string, error, int) {
	article, err := model.NewPublication(command.Title, command.Description, command.Content, command.WiterUserId, command.Type)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	message, err, status := hca.serviceCreatePublication.Run(*article)
	return message, err, status
}
