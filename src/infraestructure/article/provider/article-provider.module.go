package provider

import (
	"encoding/json"
	controller "golang-gingonic-hex-architecture/src/infraestructure/article/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/article/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/article/provider/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/response"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"
	"strconv"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/article/command"
	query "golang-gingonic-hex-architecture/src/application/article/query"
	"golang-gingonic-hex-architecture/src/application/article/query/dto"
	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/article/provider/service"
	"golang-gingonic-hex-architecture/src/infraestructure/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerInstance *controller.ControllerArticle
var once sync.Once

func ArticleProvider(conn *gorm.DB, router *gin.RouterGroup) {
	once.Do(func() {
		repositoryArticle := repository.GetRepositoryArticle(conn)
		daoArticle := dao.GetDaoArticle(conn)

		serviceRegisterarticle := infraestructureService.GetServiceCreateArticle(*repositoryArticle)

		handleCreateArticle := command.NewHandlerCreateArticle(serviceRegisterarticle)
		handleListArticles := query.NewHandlerListArticles(*daoArticle)
		handleListFiltredCompanies := query.NewHandlerListFiltredArticles(*daoArticle)

		controllerInstance = controller.NewControllerArticle(*handleCreateArticle, *handleListArticles, *handleListFiltredCompanies)
		article := router.Group("/article")

		{
			article.POST(
				"/",
				middlewares.JWTMIddleware(
					jwt.NewJWTAuthService(),
					[]string{
						middlewares.ADMINISTRATOR,
						middlewares.ARTICLE_WRITER,
					}),
				CreateAticle,
			)
			article.GET("/", ListArticles)
		}
	})
}

// Create a new article
// @Summary Create article
// @Schemes http https
// @Description Enpoint to create a article
// @Tags article
// @Accept json
// @Produce json
// @Param article body command.CommandCreateArticle true "create article"
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /article [post]
func CreateAticle(c *gin.Context) {
	var article command.CommandCreateArticle
	if err := c.ShouldBindJSON(&article); err != nil {
		response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
		return
	}
	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		response.SendError(c, err.Error(), "", http.StatusUnauthorized)
		return
	}
	article.WiterUserId = int(parsedId)

	msg, err, status := controllerInstance.Create(article)
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
// @Tags article
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /article [get]
func ListArticles(c *gin.Context) {
	query := c.Request.URL.Query()
	var data []*dto.ArticleDto
	if len(query) > 0 {
		var filters dto.FilterArticlesDto
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
