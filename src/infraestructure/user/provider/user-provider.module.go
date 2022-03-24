package provider

import (
	"golang-gingonic-hex-architecture/src/infraestructure/exceptions"
	"golang-gingonic-hex-architecture/src/infraestructure/middlewares"
	controller "golang-gingonic-hex-architecture/src/infraestructure/user/controller"
	dao "golang-gingonic-hex-architecture/src/infraestructure/user/provider/dao"
	repository "golang-gingonic-hex-architecture/src/infraestructure/user/provider/repository"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"
	"strconv"
	"sync"

	command "golang-gingonic-hex-architecture/src/application/user/command"
	query "golang-gingonic-hex-architecture/src/application/user/query"
	infraestructureService "golang-gingonic-hex-architecture/src/infraestructure/user/provider/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var controllerInstance *controller.ControllerUser
var once sync.Once

func UserProvider(conn *gorm.DB, router *gin.RouterGroup) {
	once.Do(func() {
		repositoryUser := repository.GetRepositoryUser(conn)
		daoUser := dao.GetDaoUser(conn)

		serviceRegisterUser := infraestructureService.GetServiceRegisterUser(*repositoryUser)
		serviceEditUser := infraestructureService.GetServiceEditUser(*repositoryUser)
		serviceLoginUser := infraestructureService.GetServiceLoginUser(*repositoryUser)
		serviceDeleteUser := infraestructureService.GetServiceDeleteUser(*repositoryUser)

		handleRegisterUser := command.NewHandlerRegisterUser(serviceRegisterUser)
		handleLoginUser := command.NewHandlerLoginUser(serviceLoginUser)
		handleListUsers := query.NewHandlerListUsers(*daoUser)
		handleEditUser := command.NewHandlerEditUser(serviceEditUser)
		handleDeleteUser := command.NewHandlerDeleteUser(serviceDeleteUser)

		controllerInstance = controller.NewControllerUser(*handleRegisterUser, *handleListUsers, *handleLoginUser, *handleEditUser, *handleDeleteUser)
		user := router.Group("/user")
		{
			user.POST("/", CreateUser)
			user.GET("/", ListUsers)
			user.POST("/login", Login)
			user.PATCH("/",
				middlewares.JWTMIddleware(
					jwt.NewJWTAuthService(),
					[]string{
						middlewares.ADMINISTRATOR,
						middlewares.PUBLICATION_WRITER,
						middlewares.LEATHER,
						middlewares.LEGAL_REPRESENTATIVE,
						middlewares.TECHNICAL_WORKER,
					},
				),
				Update,
			)
			user.DELETE("/",
				middlewares.JWTMIddleware(
					jwt.NewJWTAuthService(),
					[]string{
						middlewares.TECHNICAL_WORKER,
					},
				),
				Delete,
			)
		}
	})
}

// Create a new user
// @Summary Create user
// @Schemes http https
// @Description Enpoint to create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body command.CommandRegisterUser true "create user"
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user command.CommandRegisterUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}
	response := controllerInstance.Create(user)

	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})
}

// Get users
// @Summary Get users
// @Schemes http https
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /user [get]
func ListUsers(c *gin.Context) {
	data := controllerInstance.List()
	c.JSON(http.StatusOK, data)
}

// Login
// @Summary Login of user
// @Schemes http https
// @Description Enpoint to login a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body command.CommandLoginUser true "login user"
// @Success 200 {object} "hola"
// @Failture 500 {object} response.ResponseModel
// @Router /user/login [post]
func Login(c *gin.Context) {
	var credentials command.CommandLoginUser
	if err := c.ShouldBindJSON(&credentials); err != nil {
		response.SendError(c, "Invalid data: "+err.Error(), "", http.StatusBadRequest)
		return
	}
	response := controllerInstance.Login(credentials)
	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})
}

// Update a user
// @Summary Update user
// @Schemes http https
// @Description Enpoint to update a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body command.CommandEditUser true "update user"
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /user [patch]
func Update(c *gin.Context) {
	var user command.CommandEditUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}

	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}
	user.UserId = int(parsedId)

	response := controllerInstance.Update(user)

	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})
}

// Delete a user
// @Summary Delete user
// @Schemes http https
// @Description Enpoint to delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body command.CommandDeleteUser true "delete user"
// @Success 200 {object} response.ResponseModel
// @Failture 500 {object} response.ResponseModel
// @Router /user [delete]
func Delete(c *gin.Context) {
	var user command.CommandDeleteUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}

	id, _ := c.Get("id")
	parsedId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid data: "+err.Error())
		return
	}

	response := controllerInstance.Delete(int(parsedId), user)
	exceptions.ExceptionAndResponseWrapper(c, response, func() {
		c.JSON(http.StatusOK, response)
	})
}
