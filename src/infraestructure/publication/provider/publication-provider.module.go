package provider

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/infraestructure/exceptions"
	"golang-gingonic-hex-architecture/src/infraestructure/middlewares"
	controller "golang-gingonic-hex-architecture/src/infraestructure/publication/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/repository"
	"io"
	"math/rand"
	"mime/multipart"
	"os"

	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/parse"
	"net/http"
	"strconv"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/publication/command"
	query "golang-gingonic-hex-architecture/src/application/publication/query"
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"

	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
			publication.GET("/file/:id", GetFile)
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

	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid data: "+err.Error())
		return
	}

	var file multipart.File
	var header *multipart.FileHeader
	var filedirectory, fileDirectory string
	if publication.ContentType != "Texto" {
		var err error
		file, header, err = c.Request.FormFile("File")
		if err != nil {

			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		filename := header.Filename
		err = os.MkdirAll("public/", os.ModePerm)

		fileDirectory = id.(string) + strconv.Itoa(rand.Intn(10000)) + filename
		filedirectory = "public/" + fileDirectory
		out, err := os.OpenFile(filedirectory, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		publication.Content = filedirectory

	}

	if err := c.MustBindWith(&publication, binding.FormMultipart); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if publication.ContentType != "Texto" {
		publication.Content = fileDirectory
	}

	publication.WiterUserId = int(parsedId)

	response := controllerInstance.Create(publication)
	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})

}

func GetFile(c *gin.Context) {
	c.File("public/" + c.Param("id"))
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
