package infraestructure

import (
	commentProvider "golang-gingonic-hex-architecture/src/infraestructure/comment/provider"
	publicationProvider "golang-gingonic-hex-architecture/src/infraestructure/publication/provider"
	userProvider "golang-gingonic-hex-architecture/src/infraestructure/user/provider"
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
		dbProperties := DatabaseConnectionProperties{
			DATABASE_TYPE:     os.Getenv("DATABASE_TYPE"),
			DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
			DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
			DATABASE_USER:     os.Getenv("DATABASE_USER"),
			DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
			DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		}
		DATABSE_STRING_CONNECTION := CreateDatabaseStringConnetion(&dbProperties)

		conn, err := gorm.Open(postgres.Open(DATABSE_STRING_CONNECTION), &gorm.Config{
			// SkipDefaultTransaction: true,

		})
		ENV := os.Getenv("ENV")
		switch ENV {
		case "development":
			dbConnection = conn.Debug()
		case "production":
			dbConnection = conn
		}

		if err != nil {
			log.Println("Error with db connection", err)
		}

		userProvider.UserProvider(dbConnection, router)
		publicationProvider.PublicationProvider(dbConnection, router)
		commentProvider.CommentProvider(conn, router)
	})
}

type DatabaseConnectionProperties struct {
	DATABASE_TYPE,
	DATABASE_HOST,
	DATABASE_PORT,
	DATABASE_USER,
	DATABASE_PASSWORD,
	DATABASE_NAME string
}

func CreateDatabaseStringConnetion(dp *DatabaseConnectionProperties) string {
	return "host=" + dp.DATABASE_HOST +
		" user=" + dp.DATABASE_USER +
		" password=" + dp.DATABASE_PASSWORD +
		" dbname=" + dp.DATABASE_NAME +
		" port=" + dp.DATABASE_PORT +
		" sslmode=" + "disable"
}
