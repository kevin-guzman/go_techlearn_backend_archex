package dto

import "time"

type ArticleDto struct {
	Id          int
	Title       string
	Description string
	Content     string
	WiterUserId string
	WrittenAt   time.Time
}
