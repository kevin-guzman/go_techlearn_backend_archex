package repository

import "golang-gingonic-hex-architecture/src/domain/publication/model"

type RepositoryPublication interface {
	Save(publication model.Publication) error
	ExistById(id int) (bool, error)
}
