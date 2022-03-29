package dto

import (
	"golang-gingonic-hex-architecture/src/application/comment/query/dto"
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"time"

	"github.com/lib/pq"
)

type PublicationDto struct {
	Id          int
	Title       string
	Description string
	Content     string
	WiterUserId int
	WrittenAt   time.Time
	Type        model.PublicationTypes
	Categories  pq.StringArray   `gorm:"type:text[]"`
	Comments    []dto.CommentDto `gorm:"foreignKey:PublicationId"`
}
