package service

import (
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"net/http"
)

type ServiceEditUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceEditUser(UserR repository.RepositoryUser) *ServiceEditUser {
	return &ServiceEditUser{
		userRepository: UserR,
	}
}

func (seu ServiceEditUser) Run(id int, user model.User) (string, error, int) {
	LoadStringsFromService(SERVICE_EDIT)
	err := seu.userRepository.EditUser(id, user)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}
