package entity

import (
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"time"

	"golang-gingonic-hex-architecture/src/domain/publication/model"

	"gorm.io/gorm"
)

type Publication struct {
	gorm.Model
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	Content     string
	WiterUserId int
	User        entity.User `gorm:"foreignKey:WiterUserId"`
	WrittenAt   time.Time
	Type        model.PublicationTypes
}
