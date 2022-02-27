package service_test

import (
	"golang-gingonic-hex-architecture/src/domain/company/model"
	"golang-gingonic-hex-architecture/src/domain/company/service"
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/tests/utils/mocks"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	repositoryCompany mocks.MockRepositoryCompany
	cmp               = model.Company{Name: "company"}
	t                 *testing.T
)

var _ = Describe("Service create company", func() {
	BeforeSuite(func() {
		t = tReference
		errors.NewErrorCore = func(err error, trace, message string) *errors.ErrorCore {
			return &errors.ErrorCore{
				Err:     err,
				Trace:   trace,
				Message: message,
			}
		}
	})

	BeforeEach(func() {
		repositoryCompany = mocks.MockRepositoryCompany{}
	})

	It("Should create a company", func() {
		repositoryCompany.On("ExistCompanyByName", cmp.Name).Return(false, nil)
		repositoryCompany.On("Save", cmp).Return(nil)

		serviceRegisterCompanyStub := service.NewServiceRegisterCompany(&repositoryCompany)
		msg, err, code := serviceRegisterCompanyStub.Run(cmp)

		repositoryCompany.AssertNumberOfCalls(t, "ExistCompanyByName", 1)
		repositoryCompany.AssertCalled(tReference, "ExistCompanyByName", cmp.Name)
		repositoryCompany.AssertNumberOfCalls(t, "Save", 1)
		repositoryCompany.AssertCalled(tReference, "Save", cmp)
		repositoryCompany.AssertExpectations(t)

		Expect(msg).To(Equal("Company has succesfully created!"))
		Expect(err).To(BeNil())
		Expect(code).To(BeIdenticalTo(200))
	})

	It("If company already exists", func() {
		repositoryCompany.On("ExistCompanyByName", cmp.Name).Return(true, nil)

		serviceRegisterCompanyStub := service.NewServiceRegisterCompany(&repositoryCompany)
		msg, err, code := serviceRegisterCompanyStub.Run(cmp)

		repositoryCompany.AssertNumberOfCalls(t, "ExistCompanyByName", 1)
		repositoryCompany.AssertExpectations(t)

		Expect(msg).To(Equal(""))
		Expect(err.Error()).To(Equal("The company " + cmp.Name + " already exist"))
		Expect(code).To(BeIdenticalTo(500))
	})

})
