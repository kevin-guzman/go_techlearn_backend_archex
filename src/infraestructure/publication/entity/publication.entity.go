package entity

import (
	commentEntity "golang-gingonic-hex-architecture/src/infraestructure/comment/entity"
	userEntity "golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"time"

	"golang-gingonic-hex-architecture/src/domain/publication/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Publication struct {
	gorm.Model
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	Content     string
	WiterUserId int
	User        userEntity.User `gorm:"foreignKey:WiterUserId"`
	WrittenAt   time.Time
	Type        model.PublicationTypes
	Categories  pq.StringArray          `gorm:"type:text[]"`
	Comments    []commentEntity.Comment `gorm:"foreignKey:PublicationId"`
	ContentType string
}
