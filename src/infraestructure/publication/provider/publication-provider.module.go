package provider

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/infraestructure/exceptions"
	controller "golang-gingonic-hex-architecture/src/infraestructure/publication/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/repository"

	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/parse"
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
		handleListSearchedPublications := query.NewHandlerLSearchedistPublications(*daoPublication)
		handleGetOneById := query.NewHandlerGetOneById(*daoPublication)

		controllerInstance = controller.NewControllerPublication(
			*handleCreatePublication,
			*handleListPublications,
			*handleListFiltredPublications,
			*handleListSearchedPublications,
			*handleGetOneById,
		)
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
			publication.GET("/:id", GetOne)
			publication.GET("/search", SearchPublications)
		}
	})
}

// Get publications by search
// @Summary Get publications/search
// @Schemes http https
// @Description Get all publications by search
// @Tags publication
// @Accept json
// @Produce json
// @Success 200 {object} []dto.PublicationDto
// @Failture 500 {object} err
// @Router /publication/search [get]
func GetOne(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.String(http.StatusBadRequest, "El recurso con el id solicitado no existe")
	}
	var data *dto.PublicationDto
	data = controllerInstance.GetOneById(id)

	c.JSON(200, data)
}

// Get publications by search
// @Summary Get publications/search
// @Schemes http https
// @Description Get all publications by search
// @Tags publication
// @Accept json
// @Produce json
// @Success 200 {object} []dto.PublicationDto
// @Failture 500 {object} err
// @Router /publication/search [get]
func SearchPublications(c *gin.Context) {
	searchValue := c.Query("value")
	fmt.Println("Query value", searchValue)
	var data []*dto.PublicationDto
	data = controllerInstance.Search(searchValue)

	c.JSON(200, data)
}

// Create a new publication
// @Summary Create publication
// @Schemes http https
// @Description Enpoint to create a publication
// @Tags publication
// @Accept json
// @Produce json
// @Param publication body command.CommandCreatePublication true "create publication"
// @Success 200 {object} "Has succesfully created"
// @Failture 500 {object} err
// @Router /publication [post]
func CreatePublication(c *gin.Context) {
	var publication command.CommandCreatePublication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}
	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid data: "+err.Error())
		return
	}
	publication.WiterUserId = int(parsedId)

	response := controllerInstance.Create(publication)
	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})
}

// Get publications
// @Summary Get publications
// @Schemes http https
// @Description Get all publications
// @Tags publication
// @Accept json
// @Produce json
// @Success 200 {object} []dto.PublicationDto
// @Failture 500 {object} err
// @Router /publication [get]
func ListPublications(c *gin.Context) {
	query := c.Request.URL.Query()
	var data []*dto.PublicationDto
	if len(query) > 0 {
		var filters dto.FilterPublicationsDto

		if err := parse.ParseQueryToDto(query, &filters); err != nil {
			c.String(http.StatusUnauthorized, "Invalid data: "+err.Error())
			return
		}
		data = controllerInstance.ListFiltred(filters)
	} else {
		data = controllerInstance.List()
	}
	c.JSON(http.StatusOK, data)
}
