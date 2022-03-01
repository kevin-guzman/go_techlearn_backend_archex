package dto

import "golang-gingonic-hex-architecture/src/domain/article/model"

type FilterArticlesDto struct {
	Type   []model.ArticleTypes `form:"Type" json:"Type"`
	Limit  int                  `json:"Limit"`
	Offset int                  `json:"Offset"`
}
