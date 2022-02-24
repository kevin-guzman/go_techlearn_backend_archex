package service

import (
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var (
	errTraceLogin       string = "This error has ocurred in login-user.service.go"
	internalErrorLogin  string = "Internal server error"
	successMessageLogin string = "User has succesfully created!"
)

type ServiceLoginUser struct {
	userRepository repository.RepositoryUser
}

func NewServiceLoginUser(UserR repository.RepositoryUser) *ServiceLoginUser {
	return &ServiceLoginUser{
		userRepository: UserR,
	}
}

func (sru *ServiceLoginUser) Run(email, password string) (string, error, int) {
	user, err := sru.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}

	userPasswordBytes := []byte(user.Password)
	passwordBytes := []byte(password)
	err = bcrypt.CompareHashAndPassword(userPasswordBytes, passwordBytes)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Invalid credentials for user").PublicError(), http.StatusUnauthorized
	}

	token := jwt.NewJWTAuthService().GenerateToken(strconv.Itoa(user.Id), email, user.Role)
	return token, nil, http.StatusOK
}
