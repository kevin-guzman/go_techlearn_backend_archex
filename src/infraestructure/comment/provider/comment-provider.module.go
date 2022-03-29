package provider

import (
	controller "golang-gingonic-hex-architecture/src/infraestructure/comment/controller"
	repository "golang-gingonic-hex-architecture/src/infraestructure/comment/provider/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/exceptions"
	repositoryPublication "golang-gingonic-hex-architecture/src/infraestructure/publication/provider/repository"

	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"
	"strconv"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/comment/command"

	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/comment/provider/service"
	"golang-gingonic-hex-architecture/src/infraestructure/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerInstance *controller.ControllerComment
var once sync.Once

func CommentProvider(conn *gorm.DB, router *gin.RouterGroup) {
	once.Do(func() {
		repositoryPublication := repositoryPublication.GetRepositoryPublication(conn)
		repositoryComment := repository.GetRepositoryComment(conn)

		serviceCreateComment := infraestructureService.GetServiceCreateComment(*repositoryComment, *repositoryPublication)

		handleCreateComment := command.NewHandlerCreateComment(serviceCreateComment)

		controllerInstance = controller.NewControllerComment(*handleCreateComment)
		comment := router.Group("/comment")

		{
			comment.POST(
				"/",
				middlewares.JWTMIddleware(
					jwt.NewJWTAuthService(),
					[]string{
						middlewares.ADMINISTRATOR,
						middlewares.PUBLICATION_WRITER,
						middlewares.TECHNICAL_WORKER,
						middlewares.LEATHER,
					}),
				CreateComment,
			)
		}
	})
}

// Create a new comment
// @Summary Create comment
// @Schemes http https
// @Description Enpoint to create a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body command.CommandCreateComment true "create comment"
// @Success 200 {object} "Has succesfully created"
// @Failture 500 {object} err
// @Router /comment [post]
func CreateComment(c *gin.Context) {
	var comment command.CommandCreateComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}
	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		c.String(http.StatusUnauthorized, "Invalid data: "+err.Error())
		return
	}
	comment.UserId = int(parsedId)

	response := controllerInstance.Create(comment)
	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})
}
