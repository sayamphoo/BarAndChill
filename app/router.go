package app

import (
	"os"
	"path/filepath"
	"sayamphoo/microservice/controllers"
	"sayamphoo/microservice/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var memberController *controllers.MemberController
var songController *controllers.SongController
var tableController *controllers.TableController

func InitializeController() {

	memberController = &controllers.MemberController{}
	songController = &controllers.SongController{}
	tableController = &controllers.TableController{}

	router := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AddAllowHeaders("Authorization")

	router.Use(cors.New(config))
	router.Use(middleware.ExceptionHandling())
	router.Use(middleware.RequestMiddleware())

	//Router -> Controller
	customRouter(router)
	ownerRouter(router)

	//path get picture (Binary)
	router.GET("/image/:filename", func(ctx *gin.Context) {
		filename := ctx.Param("filename")

		filepath := filepath.Join("./resource", filename)
		_, err := os.Stat(filepath)

		if err != nil {
			return
		}
		ctx.File(filepath)
	})

	router.Run(os.Getenv("SERVER_PORT"))
}
