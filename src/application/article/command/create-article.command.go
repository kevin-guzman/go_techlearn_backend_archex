package command

import "golang-gingonic-hex-architecture/src/domain/article/model"

type CommandCreateArticle struct {
	Title       string             `json:"Title" binding:"required,min=5"`
	Description string             `json:"Description" binding:"required,min=20"`
	Content     string             `json:"Content" binding:"required,min=90"`
	WiterUserId int                `json:"WiterUserId"`
	Type        model.ArticleTypes `json:"Type" binding:"required"`
}
