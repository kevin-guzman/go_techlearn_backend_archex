package repository

import "golang-gingonic-hex-architecture/src/domain/user/model"

type RepositoryUser interface {
	ExistUserNameAndEmail(name, email string) (bool, error)
	Save(user model.User) error
	GetUserByEmail(email string) (model.User, error)
	EditUser(id int, newUser model.User) error
	Delete(id int) error
}
