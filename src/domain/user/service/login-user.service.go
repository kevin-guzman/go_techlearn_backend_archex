package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
		return "", fmt.Errorf("The user doesnt exist"), http.StatusInternalServerError
	}

	userPasswordBytes := []byte(user.Password)
	passwordBytes := []byte(password)
	err = bcrypt.CompareHashAndPassword(userPasswordBytes, passwordBytes)
	if err != nil {
		return "", fmt.Errorf("Invalid credentials for user"), http.StatusUnauthorized
	}

	token := jwt.NewJWTAuthService().GenerateToken(user.Id, email, user.Role)
	return token, nil, http.StatusOK
}
