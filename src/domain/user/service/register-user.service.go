package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
)

type ServiceRegisterUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceRegisterUser(UserR repository.RepositoryUser) *ServiceRegisterUser {
	return &ServiceRegisterUser{
		userRepository: UserR,
	}
}

func (sru *ServiceRegisterUser) Run(user model.User) interface{} {
	LoadStringsFromService(SERVICE_REGISTER)
	existUserName, err := sru.userRepository.ExistUserNameAndEmail(user.Name, user.Email)
	if err != nil {
		return errors.NewErrorPort(err)
	}
	if existUserName {
		err := fmt.Errorf("El usuario con nombre %s o email %s ya existe", user.Name, user.Email)
		return errors.NewErrorUserAlreadyExist(err, err.Error())
	}

	err = sru.userRepository.Save(user)
	if err != nil {
		fmt.Println(errTrace)
		return errors.NewErrorPort(err)
	}

	return successMessage
}
