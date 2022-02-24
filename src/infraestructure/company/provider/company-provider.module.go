package provider

import (
	controller "golang-gingonic-hex-architecture/src/infraestructure/company/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/company/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/company/provider/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/response"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/company/command"
	query "golang-gingonic-hex-architecture/src/application/company/query"
	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/company/provider/service"
	"golang-gingonic-hex-architecture/src/infraestructure/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerInstance *controller.ControllerCompany
var once sync.Once

func CompanyProvider(conn *gorm.DB, router *gin.RouterGroup) {
	once.Do(func() {
		repositoryCompany := repository.GetRepositoryCompany(conn)
		daoCompany := dao.GetDaoCompany(conn)

		serviceRegistercompany := infraestructureService.GetServiceRegisterCompany(*repositoryCompany)

		handleRegistercompany := command.NewHandlerRegisterCompany(serviceRegistercompany)
		handleListCompanies := query.NewHandlerListCompanies(*daoCompany)

		controllerInstance = controller.NewControllerCompany(*handleRegistercompany, *handleListCompanies)
		company := router.Group("/company")

		{
			company.POST(
				"/",
				middlewares.JWTMIddleware(
					jwt.NewJWTAuthService(),
					[]string{
						middlewares.ADMINISTRATOR,
						middlewares.LEGAL_REPRESENTATIVE,
					}),
				CreateCompany,
			)
			company.GET("/", ListCompanies)
		}
	})
}

// Create a new company
// @Summary Create company
// @Schemes http https
// @Description Enpoint to create a company
// @Tags company
// @Accept json
// @Produce json
// @Param company body command.CommandRegisterCompany true "create company"
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /company [post]
func CreateCompany(c *gin.Context) {
	var company command.CommandRegisterCompany
	if err := c.ShouldBindJSON(&company); err != nil {
		response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
		return
	}
	msg, err, status := controllerInstance.Create(company)
	if err != nil {
		response.SendError(c, err.Error(), msg, status)
		return
	}
	response.SendSucess(c, msg, status, nil)
}

// Get companies
// @Summary Get companies
// @Schemes http https
// @Description Get all companies
// @Tags company
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /company [get]
func ListCompanies(c *gin.Context) {
	data := controllerInstance.List()
	response.SendSucess(c, "", 200, data)
}
