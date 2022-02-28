package command

import (
	"golang-gingonic-hex-architecture/src/domain/article/model"
	"golang-gingonic-hex-architecture/src/domain/article/service"
	"net/http"
)

type HandlerCreateArticle struct {
	serviceCreateArticle service.ServiceCreateArticle
}

func NewHandlerCreateArticle(sca *service.ServiceCreateArticle) *HandlerCreateArticle {
	return &HandlerCreateArticle{
		serviceCreateArticle: *sca,
	}
}

func (hca *HandlerCreateArticle) Run(command CommandCreateArticle) (string, error, int) {
	article, err := model.NewArticle(command.Title, command.Description, command.Content, command.WiterUserId)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	message, err, status := hca.serviceCreateArticle.Run(*article)
	return message, err, status
}
