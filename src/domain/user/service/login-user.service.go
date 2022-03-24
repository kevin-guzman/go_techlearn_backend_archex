package service

import (
	userDto "golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"strconv"

	dtoUtil "github.com/dranikpg/dto-mapper"
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

func (sru *ServiceLoginUser) Run(email, password string) interface{} {
	LoadStringsFromService(SERVICE_LOGIN)
	user, err := sru.userRepository.GetUserByEmail(email)
	if err != nil {
		return errors.NewErrorPort(err)
	}

	userPasswordBytes := []byte(user.Password)
	passwordBytes := []byte(password)
	err = bcrypt.CompareHashAndPassword(userPasswordBytes, passwordBytes)
	if err != nil {
		return errors.NewErrorUserCredentials(err)
	}

	token := jwt.NewJWTAuthService().GenerateToken(strconv.Itoa(user.Id), email, user.Role)
	var _user userDto.UserDto
	dtoUtil.Map(&_user, user)
	data := map[string]interface{}{"token": token, "user": _user}

	return data
}
