package repository

import "golang-gingonic-hex-architecture/src/domain/article/model"

type RepositoryArticle interface {
	Save(article model.Article) error
}
