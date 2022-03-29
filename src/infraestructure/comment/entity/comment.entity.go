package entity

import (
	// publicationEntity "golang-gingonic-hex-architecture/src/infraestructure/publication/entity"
	userEntity "golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id             int `gorm:"primaryKey"`
	CommentTitle   string
	CommentContent string
	UserId         int
	User           userEntity.User `gorm:"foreignKey:UserId"`
	PublicationId  int
	// Publication   publicationEntity.Publication `gorm:"foreignKey:PublicationId"`
	WrittenAt time.Time
}
