package model

import (
	"time"
)

type Comment struct {
	Id            int
	Title         string
	Content       string
	UserId        int
	WrittenAt     time.Time
	PublicationId int
}

func NewComment(tittle, content string, userId, publicationId int) *Comment {
	return &Comment{
		Title:         tittle,
		Content:       content,
		UserId:        userId,
		PublicationId: publicationId,
		WrittenAt:     time.Now(),
	}
}
