package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	doc "golang-gingonic-hex-architecture/src/docs"

	"golang-gingonic-hex-architecture/src/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	// godotenv.Load("../.env")
	// PORT := os.Getenv("PORT")
	// CONTEXT_PATH := os.Getenv("CONTEXT_PATH")
	run_env := os.Getenv("ENVIRONMENT")
	path_env := "./env/development.env"
	switch run_env {
	case "test":
		path_env = "./env/testing.env"
	case "production":
		path_env = "./env/production.env"
	default:
		path_env = "./env/development.env"
	}

	if err := godotenv.Load(path_env); err != nil {
		log.Fatal("Error reading .env file\n", err)
	}

	ENV := os.Getenv("ENV")
	switch ENV {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	}

	PORT := os.Getenv("APPLICATION_PORT")
	CONTEXT_PATH := os.Getenv("CONTEXT_PATH")

	doc.SwaggerInfo_swagger.BasePath = "/" + CONTEXT_PATH
	server := gin.Default()
	path := server.Group(CONTEXT_PATH)
	{
		infraestructure.InitInfraestructure(path)
		url := ginSwagger.URL("http://localhost" + PORT + "/" + CONTEXT_PATH + "/swagger/doc.json")
		path.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	if err := server.Run(PORT); err != nil {
		log.Fatal(err)
	}

}
