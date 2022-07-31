package app

import (
	"golte/app/config"
	"golte/app/middlewares"
	"golte/app/repository"
	"golte/app/routes"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogOutPut() error {
	f, err := os.Create("./gin.log")
	if err != nil {
		return err
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return nil
}

func SetupRouter() *gin.Engine {
	logger := setupLogOutPut()
	gin.SetMode(gin.ReleaseMode)

	if logger != nil {
		log.Fatalf("Failed create logger. Error : %v", logger.Error())
	}

	router := gin.Default()

	router.Static("static", "./templates/static")

	router.LoadHTMLGlob("./templates/*.html")

	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.Auth)

	routes.Route(router)

	return router
}

func Run(address string) {
	db := config.ConnectDB()
	if db != nil {
		log.Fatal(db.Error())
	}
	repository.RunMigration(config.DB)
	router := SetupRouter()
	err := router.Run(address)
	if err != nil {
		log.Fatalf("Failed listen at %v. Error : %v", address, err.Error())
	}
}
