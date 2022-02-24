package mocks

import (
	"golang-gingonic-hex-architecture/src/domain/company/model"
	"golang-gingonic-hex-architecture/src/infraestructure/company/adaptor/repository"

	"github.com/stretchr/testify/mock"
)

type MockRepositoryCompany struct {
	repository.RepositoryCompanyPostgreSql
	mock.Mock
}

func (m *MockRepositoryCompany) ExistCompanyByName(name string) (bool, error) {
	args := m.Called(name)
	var returned bool
	if len(args) > 1 {
		returned = args.Get(0).(bool)
	}
	return returned, nil
}

func (m *MockRepositoryCompany) ExistCompanyById(id int) (bool, error) {
	args := m.Called(id)
	var returned bool
	if len(args) > 1 {
		returned = args.Get(0).(bool)
	}
	return returned, nil
}

func (m *MockRepositoryCompany) Save(c model.Company) error {
	args := m.Called(c)
	var returned error
	if len(args) > 1 {
		returned = args.Get(0).(error)
	}
	return returned
}
