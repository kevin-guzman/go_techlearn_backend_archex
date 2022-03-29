package dto

import (
	"time"
)

type CommentDto struct {
	Id             int
	CommentTitle   string
	CommentContent string
	UserId         int
	WrittenAt      time.Time
	PublicationId  int
}
