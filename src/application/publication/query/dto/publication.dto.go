package dto

import (
	"golang-gingonic-hex-architecture/src/domain/publication/model"
	"time"
)

type PublicationDto struct {
	Id          int
	Title       string
	Description string
	Content     string
	WiterUserId string
	WrittenAt   time.Time
	Type        model.PublicationTypes
}
