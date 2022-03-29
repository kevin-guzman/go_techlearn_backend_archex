package controller

import (
	"golang-gingonic-hex-architecture/src/application/comment/command"
)

type ControllerComment struct {
	handlerCreateComment command.HandlerCreateComment
}

func NewControllerComment(
	hcc command.HandlerCreateComment,
) *ControllerComment {
	return &ControllerComment{
		handlerCreateComment: hcc,
	}
}

func (cp *ControllerComment) Create(command command.CommandCreateComment) interface{} {
	return cp.handlerCreateComment.Run(command)
}
