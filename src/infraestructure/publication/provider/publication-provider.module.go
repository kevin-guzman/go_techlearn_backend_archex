package provider

import (
	"encoding/json"
	controller "golang-gingonic-hex-architecture/src/infraestructure/publication/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/response"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"
	"strconv"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/publication/command"
	query "golang-gingonic-hex-architecture/src/application/publication/query"
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/middlewares"
	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerInstance *controller.ControllerPublication
var once sync.Once

func PublicationProvider(conn *gorm.DB, router *gin.RouterGroup) {
	once.Do(func() {
		repositoryPublication := repository.GetRepositoryPublication(conn)
		daoPublication := dao.GetDaoPublication(conn)

		serviceRegisterPublication := infraestructureService.GetServiceCreatePublication(*repositoryPublication)

		handleCreatePublication := command.NewHandlerCreatePublication(serviceRegisterPublication)
		handleListPublications := query.NewHandlerListPublication(*daoPublication)
		handleListFiltredPublications := query.NewHandlerListFiltredPublication(*daoPublication)

		controllerInstance = controller.NewControllerPublication(*handleCreatePublication, *handleListPublications, *handleListFiltredPublications)
		publication := router.Group("/publication")

		{
			publication.POST(
				"/",
				middlewares.JWTMIddleware(
					jwt.NewJWTAuthService(),
					[]string{
						middlewares.ADMINISTRATOR,
						middlewares.PUBLICATION_WRITER,
					}),
				CreatePublication,
			)
			publication.GET("/", ListPublications)
		}
	})
}

// Create a new publication
// @Summary Create publication
// @Schemes http https
// @Description Enpoint to create a publication
// @Tags publication
// @Accept json
// @Produce json
// @Param publication body command.CommandCreatePublication true "create publication"
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /publication [post]
func CreatePublication(c *gin.Context) {
	var publication command.CommandCreatePublication
	if err := c.ShouldBindJSON(&publication); err != nil {
		response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
		return
	}
	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		response.SendError(c, err.Error(), "", http.StatusUnauthorized)
		return
	}
	publication.WiterUserId = int(parsedId)

	msg, err, status := controllerInstance.Create(publication)
	if err != nil {
		response.SendError(c, err.Error(), msg, status)
		return
	}
	response.SendSucess(c, msg, status, nil)
}

// Get publications
// @Summary Get publications
// @Schemes http https
// @Description Get all publications
// @Tags publication
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /publication [get]
func ListPublications(c *gin.Context) {
	query := c.Request.URL.Query()
	var data []*dto.PublicationDto
	if len(query) > 0 {
		var filters dto.FilterPublicationsDto
		if bytes, err := json.Marshal(query); err != nil {
			response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
			return
		} else {
			if err = json.Unmarshal(bytes, &filters); err != nil {
				response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
				return
			}
			data = controllerInstance.ListFiltred(filters)
		}
	} else {
		data = controllerInstance.List()
	}
	response.SendSucess(c, "", 200, data)
}
