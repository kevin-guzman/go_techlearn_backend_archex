package infraestructure

import (
	"golang-gingonic-hex-architecture/src/infraestructure/user/provider"
	"os"
	"sync"

	"github.com/gin-gonic/gin"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB
var once sync.Once

var InitInfraestructure = func(router *gin.RouterGroup) {

	once.Do(func() {
		DATABSE_STRING_CONNECTION := os.Getenv("DATABSE_STRING_CONNECTION")
		conn, err := gorm.Open(postgres.Open(DATABSE_STRING_CONNECTION), &gorm.Config{})
		dbConnection = conn //.Debug()

		if err != nil {
			log.Println("Error with db connection", err)
		}

		provider.UserProvider(dbConnection, router)
	})
}
