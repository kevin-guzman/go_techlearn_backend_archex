package dto

import "golang-gingonic-hex-architecture/src/domain/publication/model"

type FilterPublicationsDto struct {
	Type   []model.PublicationTypes `form:"Type" json:"Type"`
	Limit  int                      `json:"Limit"`
	Offset int                      `json:"Offset"`
}
