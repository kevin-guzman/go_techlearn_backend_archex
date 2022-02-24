package service_test

import (
	"golang-gingonic-hex-architecture/src/domain/errors"
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/domain/user/service"
	"golang-gingonic-hex-architecture/tests/utils/mocks"
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	repositoryUser    mocks.MockRepositoryUser
	repositoryCompany mocks.MockRepositoryCompany
	usr               = model.User{Name: "Juan", CompanyId: 1}
	t                 *testing.T
)

var _ = Describe("Service create user", func() {
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
		repositoryUser = mocks.MockRepositoryUser{}
		repositoryCompany = mocks.MockRepositoryCompany{}
	})

	It("Should create a user", func() {
		repositoryUser.On("ExistUserName", usr.Name).Return(false, nil)
		repositoryUser.On("Save", usr).Return(nil)
		repositoryCompany.On("ExistCompanyById", usr.CompanyId).Return(true, nil)

		serviceRegisterUserStub := service.NewServiceRegisterUser(&repositoryUser, &repositoryCompany)
		msg, err, code := serviceRegisterUserStub.Run(usr)

		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
		repositoryUser.AssertCalled(t, "ExistUserName", usr.Name)
		repositoryUser.AssertNumberOfCalls(t, "Save", 1)
		repositoryUser.AssertCalled(t, "Save", usr)
		repositoryUser.AssertExpectations(t)
		repositoryCompany.AssertNumberOfCalls(t, "ExistCompanyById", 1)
		repositoryCompany.AssertCalled(tReference, "ExistCompanyById", usr.CompanyId)
		repositoryCompany.AssertExpectations(t)

		Expect(msg).To(Equal("User has succesfully created!"))
		Expect(err).To(BeNil())
		Expect(code).To(BeIdenticalTo(200))
	})

	It("If user already exists", func() {
		repositoryUser.On("ExistUserName", usr.Name).Return(true, nil)

		serviceRegisterUserStub := service.NewServiceRegisterUser(&repositoryUser, &repositoryCompany)
		msg, err, code := serviceRegisterUserStub.Run(usr)

		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
		repositoryUser.AssertCalled(t, "ExistUserName", usr.Name)
		repositoryUser.AssertNumberOfCalls(t, "Save", 0)
		repositoryUser.AssertExpectations(t)
		repositoryCompany.AssertNumberOfCalls(t, "ExistCompanyById", 0)
		repositoryCompany.AssertExpectations(t)

		Expect(msg).To(Equal(""))
		Expect(err.Error()).To(Equal("The username " + usr.Name + " already exist"))
		Expect(code).To(BeIdenticalTo(500))
	})

	It("If the company doesnt exist", func() {
		repositoryUser.On("ExistUserName", usr.Name).Return(false, nil)
		repositoryCompany.On("ExistCompanyById", usr.CompanyId).Return(false, nil)

		serviceRegisterUserStub := service.NewServiceRegisterUser(&repositoryUser, &repositoryCompany)
		msg, err, code := serviceRegisterUserStub.Run(usr)

		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
		repositoryUser.AssertCalled(t, "ExistUserName", usr.Name)
		repositoryUser.AssertNumberOfCalls(t, "Save", 0)
		repositoryUser.AssertExpectations(t)
		repositoryCompany.AssertNumberOfCalls(t, "ExistCompanyById", 1)
		repositoryCompany.AssertCalled(t, "ExistCompanyById", usr.CompanyId)
		repositoryCompany.AssertExpectations(t)

		Expect(msg).To(Equal(""))
		Expect(err.Error()).To(Equal("The company with id " + strconv.Itoa(usr.CompanyId) + " doesnt exist"))
		Expect(code).To(BeIdenticalTo(500))
	})

})
