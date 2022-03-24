package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"

	"golang.org/x/crypto/bcrypt"
)

type ServiceDeleteUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceDeleteUser(UserR repository.RepositoryUser) *ServiceDeleteUser {
	return &ServiceDeleteUser{
		userRepository: UserR,
	}
}

func (sdu *ServiceDeleteUser) Run(id int, user model.User) interface{} {
	LoadStringsFromService(SERVICE_DELETE)
	currentUser, err := sdu.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		errMsg := fmt.Errorf("El usuario con email %s no existe", user.Email)
		return errors.NewErrorUserDoentExist(err, errMsg.Error())
	}

	if id != currentUser.Id {
		errMsg := fmt.Errorf("No tienes acceso para borrar el usuario")
		return errors.NewErrorUserPermission(errMsg, errMsg.Error())
	}

	userPasswordBytes := []byte(currentUser.Password)
	passwordBytes := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(userPasswordBytes, passwordBytes)
	if err != nil {
		return errors.NewErrorUserCredentials(err)
	}

	err = sdu.userRepository.Delete(id)
	if err != nil {
		return errors.NewErrorPort(err)
	}

	return successMessage
}
