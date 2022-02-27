package controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"golang-gingonic-hex-architecture/src/application/user/command"
	"golang-gingonic-hex-architecture/src/application/user/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure"
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider"
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider/dao"

	repositoryCompanyProvider "golang-gingonic-hex-architecture/src/infraestructure/company/provider/repository"
	repositoryUserProvider "golang-gingonic-hex-architecture/src/infraestructure/user/provider/repository"

	"golang-gingonic-hex-architecture/tests/utils"
	"golang-gingonic-hex-architecture/tests/utils/mocks"
	"log"

	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	interfaceRepositoryCompany "golang-gingonic-hex-architecture/src/domain/company/port/repository"
	"golang-gingonic-hex-architecture/src/domain/errors"
	interfaceDao "golang-gingonic-hex-architecture/src/domain/user/port/dao"
	interfaceRepositoryUser "golang-gingonic-hex-architecture/src/domain/user/port/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	repositoryUser    mocks.MockRepositoryUser
	repositoryCompany mocks.MockRepositoryCompany
	daoUser           mocks.MockDaoUser
	server            *gin.Engine
	context           *gin.Context
	t                 *testing.T
	expected, got     gin.H
)

var _ = Describe("Tests of the user controller", func() {
	BeforeEach(func() {
		repositoryUser = mocks.MockRepositoryUser{}
		repositoryCompany = mocks.MockRepositoryCompany{}
		daoUser = mocks.MockDaoUser{}
	})

	AfterEach(func() {
		context.Done()
	})

	BeforeSuite(func() {
		t = tReference
		errors.NewErrorCore = func(err error, trace, message string) *errors.ErrorCore {
			return &errors.ErrorCore{
				Err:     err,
				Trace:   trace,
				Message: message,
			}
		}
		dao.GetDaoUser = func(conn *gorm.DB) *interfaceDao.DaoUser {
			ru := &daoUser
			iru := interfaceDao.DaoUser(ru)
			_, _ = iru, ru
			return &iru
		}
		repositoryUserProvider.GetRepositoryUser = func(conn *gorm.DB) *interfaceRepositoryUser.RepositoryUser {
			ru := &repositoryUser
			iru := interfaceRepositoryUser.RepositoryUser(ru)
			_, _ = iru, ru
			return &iru
		}
		repositoryCompanyProvider.GetRepositoryCompany = func(conn *gorm.DB) *interfaceRepositoryCompany.RepositoryCompany {
			ru := &repositoryCompany
			iru := interfaceRepositoryCompany.RepositoryCompany(ru)
			_, _ = iru, ru
			return &iru
		}
		infraestructure.InitInfraestructure = func(router *gin.RouterGroup) {
			/* Mocking the databse connection to do it as a null or mocked connection */
			var db *sql.DB
			db, _, _ = sqlmock.New()
			dial := postgres.New(postgres.Config{
				DriverName: "postgres",
				Conn:       db,
			})
			conn, _ := gorm.Open(dial, &gorm.Config{
				SkipDefaultTransaction: true,
			})
			provider.UserProvider(conn, router)
		}
		gin.SetMode(gin.TestMode)
		context, server = gin.CreateTestContext(httptest.NewRecorder())
		path_env := "../../../../env/testing.env"
		if err := godotenv.Load(path_env); err != nil {
			log.Fatal("Error reading .env file\n", err)
		}
		CONTEXT_PATH := os.Getenv("CONTEXT_PATH")
		path := server.Group(CONTEXT_PATH)
		infraestructure.InitInfraestructure(path)
	})

	It("Should list the registered users", func() {
		users := []*dto.UserDto{{Name: "name", Creation_date: time.Now()}}
		utils.JSONParse(
			gin.H{
				"data":    users,
				"status":  200,
				"error":   "",
				"message": "",
			},
			&expected,
		)
		daoUser.On("List").Return(users)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/user/", nil)
		server.ServeHTTP(w, req)

		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			log.Fatal("Error on unmarsall reponse ", err)
		}

		Expect(w.Code).To(BeNumerically("==", http.StatusOK))
		Expect(expected).To(Equal(got))

		daoUser.AssertCalled(t, "List")
		daoUser.AssertNumberOfCalls(t, "List", 1)
	})

	It("Should fail on create a user with a short password", func() {
		user := command.CommandRegisterUser{Name: "Juan Jose", Password: "12", Email: "emailoftest", Role: "role", CompanyId: 12}
		userBytes, _ := json.Marshal(user)
		const message string = "Invalid data: Key: 'CommandRegisterUser.Password' Error:Field validation for 'Password' failed on the 'min' tag"
		utils.JSONParse(
			gin.H{
				"data":    nil,
				"status":  400,
				"error":   message,
				"message": "",
			},
			&expected,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/", bytes.NewBuffer(userBytes))
		server.ServeHTTP(w, req)

		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			log.Fatal("Error on unmarsall reponse", err)
		}

		Expect(w.Code).To(BeNumerically("==", http.StatusBadRequest))
		Expect(expected).To(Equal(got))
		Expect(got["error"]).To(Equal(message))
	})

	It("Should fail because the user has already exist", func() {
		user := command.CommandRegisterUser{Name: "Juan Jose", Password: "12267gdgweg3", Email: "emailoftest", Role: "role", CompanyId: 12}
		userBytes, _ := json.Marshal(user)
		var message string = "The username " + user.Name + " already exist"
		utils.JSONParse(
			gin.H{
				"data":    nil,
				"status":  500,
				"error":   message,
				"message": "",
			},
			&expected,
		)
		repositoryUser.On("ExistUserName", user.Name).Return(true, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/", bytes.NewBuffer(userBytes))
		server.ServeHTTP(w, req)

		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			log.Fatal("Error on unmarsall reponse ex", err)
		}

		Expect(w.Code).To(BeNumerically("==", http.StatusInternalServerError))
		Expect(expected).To(Equal(got))
		Expect(got["error"]).To(Equal(message))

		repositoryUser.AssertCalled(t, "ExistUserName", user.Name)
		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
	})

	It("Should fail because the company doesn't exist", func() {
		user := command.CommandRegisterUser{Name: "Juan Jose", Password: "12267gdgweg3", Email: "emailoftest", Role: "role", CompanyId: 12}
		userBytes, _ := json.Marshal(user)
		var message string = fmt.Sprintf("The company with id %d doesnt exist", user.CompanyId)
		utils.JSONParse(
			gin.H{
				"data":    nil,
				"status":  500,
				"error":   message,
				"message": "",
			},
			&expected,
		)

		repositoryUser.On("ExistUserName", user.Name).Return(false, nil)
		repositoryCompany.On("ExistCompanyById", user.CompanyId).Return(false, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/", bytes.NewBuffer(userBytes))
		server.ServeHTTP(w, req)

		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			log.Fatal("Error on unmarsall reponse ex", err)
		}

		Expect(w.Code).To(BeNumerically("==", http.StatusInternalServerError))
		Expect(expected).To(Equal(got))
		Expect(got["error"]).To(Equal(message))

		repositoryUser.AssertCalled(t, "ExistUserName", user.Name)
		repositoryUser.AssertNumberOfCalls(t, "ExistUserName", 1)
	})

})
