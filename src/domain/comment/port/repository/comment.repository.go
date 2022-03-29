package repository

import "golang-gingonic-hex-architecture/src/domain/comment/model"

type RepositoryComment interface {
	Save(comment model.Comment) error
}
