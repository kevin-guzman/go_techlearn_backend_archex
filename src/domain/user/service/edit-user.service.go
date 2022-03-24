package service

import (
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
)

type ServiceEditUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceEditUser(UserR repository.RepositoryUser) *ServiceEditUser {
	return &ServiceEditUser{
		userRepository: UserR,
	}
}

func (seu ServiceEditUser) Run(id int, user model.User) interface{} {
	LoadStringsFromService(SERVICE_EDIT)
	err := seu.userRepository.EditUser(id, user)
	if err != nil {
		return errors.NewErrorPort(err)
	}

	return successMessage
}
