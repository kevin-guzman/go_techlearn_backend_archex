package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"net/http"

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

func (sdu *ServiceDeleteUser) Run(id int, user model.User) (string, error, int) {
	LoadStringsFromService(SERVICE_DELETE)
	currentUser, err := sdu.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		errMsg := fmt.Errorf("The user with email %s doesn't exist", user.Email)
		return "", errors.NewErrorCore(err, errTrace, errMsg.Error()).PublicError(), http.StatusInternalServerError
	}

	if id != currentUser.Id {
		errMsg := fmt.Errorf("TYou don't have acces to delete this user")
		return "", errors.NewErrorCore(errMsg, errTrace, errMsg.Error()).PublicError(), http.StatusInternalServerError
	}

	userPasswordBytes := []byte(currentUser.Password)
	passwordBytes := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(userPasswordBytes, passwordBytes)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Invalid credentials for user").PublicError(), http.StatusUnauthorized
	}

	err = sdu.userRepository.Delete(id)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}
