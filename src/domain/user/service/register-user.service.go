package service

import (
	"fmt"
	companyRepository "golang-gingonic-hex-architecture/src/domain/company/port/repository"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"net/http"
)

var (
	errTrace       string = "This error has ocurred in register-user.service.go"
	internalError  string = "Internal server error"
	successMessage string = "User has succesfully created!"
)

type ServiceRegisterUser struct {
	userRepository    repository.RepositoryUser
	companyRepository companyRepository.RepositoryCompany
}

func NewServiceRegisterUser(UserR repository.RepositoryUser, CompR companyRepository.RepositoryCompany) *ServiceRegisterUser {
	return &ServiceRegisterUser{
		userRepository:    UserR,
		companyRepository: CompR,
	}
}

func (sru *ServiceRegisterUser) Run(user model.User) (string, error, int) {
	existUserName, err := sru.userRepository.ExistUserName(user.Name)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}
	if existUserName {
		err := fmt.Errorf("The username %s already exist", user.Name)
		return "", errors.NewErrorCore(err, errTrace, err.Error()).PublicError(), http.StatusInternalServerError
	}

	existCompany, err := sru.companyRepository.ExistCompanyById(user.CompanyId)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}
	if !existCompany {
		err := fmt.Errorf("The company with id %d doesnt exist", user.CompanyId)
		return "", errors.NewErrorCore(err, errTrace, err.Error()).PublicError(), http.StatusInternalServerError
	}

	err = sru.userRepository.Save(user)
	if err != nil {
		fmt.Println(errTrace)
		return "", errors.NewErrorCore(err, errTrace, internalError).PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}
