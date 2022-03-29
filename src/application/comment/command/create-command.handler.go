package command

import (
	"golang-gingonic-hex-architecture/src/domain/comment/model"
	"golang-gingonic-hex-architecture/src/domain/comment/service"
)

type HandlerCreateComment struct {
	serviceCreatePublication service.ServiceCreateComment
}

func NewHandlerCreateComment(scc *service.ServiceCreateComment) *HandlerCreateComment {
	return &HandlerCreateComment{
		serviceCreatePublication: *scc,
	}
}

func (hca *HandlerCreateComment) Run(command CommandCreateComment) interface{} {
	comment := model.NewComment(command.Title, command.Content, command.UserId, command.PublicationId)

	return hca.serviceCreatePublication.Run(*comment)
}
