package service

import (
	"fmt"
	companyRepository "golang-gingonic-hex-architecture/src/domain/company/port/repository"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/port/repository"
	"net/http"
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
	LoadStringsFromService(SERVICE_REGISTER)
	existUserName, err := sru.userRepository.ExistUserNameAndEmail(user.Name, user.Email)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}
	if existUserName {
		err := fmt.Errorf("The user with name %s or email %s already exist", user.Name, user.Email)
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
