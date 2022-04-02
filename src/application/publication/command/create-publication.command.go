package command

import (
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"mime/multipart"

	"github.com/lib/pq"
)

type CommandCreatePublication struct {
	Title       string                 `form:"Title" binding:"required,min=5"`
	Description string                 `form:"Description" binding:"required,min=20"`
	Content     string                 `form:"Content"`
	Type        model.PublicationTypes `form:"Type" binding:"required"`
	Categories  pq.StringArray         `form:"Categories"`
	ContentType string                 `form:"ContentType" binding:"required"`
	File        *multipart.FileHeader  `form:"File"`
	WiterUserId int
}
