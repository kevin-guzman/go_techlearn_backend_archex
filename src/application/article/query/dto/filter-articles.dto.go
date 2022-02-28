package dto

import "golang-gingonic-hex-architecture/src/domain/article/model"

type FilterArticlesDto struct {
	Type []model.ArticleTypes `json:"Type" binding:"min=5"`
}
