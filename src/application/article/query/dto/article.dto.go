package dto

import (
	"golang-gingonic-hex-architecture/src/domain/article/model"
	"time"
)

type ArticleDto struct {
	Id          int
	Title       string
	Description string
	Content     string
	WiterUserId string
	WrittenAt   time.Time
	Type        model.ArticleTypes
}
