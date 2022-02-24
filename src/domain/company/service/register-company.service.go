package service

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/domain/company/model"
	"golang-gingonic-hex-architecture/src/domain/company/port/repository"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"net/http"
)

var (
	errTrace       string = "This error has ocurred in create-company.service.go"
	internalError  string = "Internal server error"
	successMessage string = "Company has succesfully created!"
)

type ServiceRegisterCompany struct {
	companyRepository repository.RepositoryCompany
}

func NewServiceRegisterCompany(CompR repository.RepositoryCompany) *ServiceRegisterCompany {
	return &ServiceRegisterCompany{
		companyRepository: CompR,
	}
}

func (src *ServiceRegisterCompany) Run(company model.Company) (string, error, int) {
	existCompany, err := src.companyRepository.ExistCompanyByName(company.Name)
	if err != nil {
		return "", errors.NewErrorCore(err, errTrace, "Service error").PublicError(), http.StatusInternalServerError
	}
	if existCompany {
		err := fmt.Errorf("The company %s already exist", company.Name)
		return "", errors.NewErrorCore(err, errTrace, err.Error()).PublicError(), http.StatusInternalServerError
	}

	err = src.companyRepository.Save(company)
	if err != nil {
		fmt.Println(errTrace)
		return "", errors.NewErrorCore(err, errTrace, internalError).PublicError(), http.StatusInternalServerError
	}

	return successMessage, nil, http.StatusOK
}
