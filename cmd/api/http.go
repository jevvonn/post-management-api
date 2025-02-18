package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jevvonn/post-management-api/database"

	_ "github.com/jevvonn/post-management-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	postDelivery "github.com/jevvonn/post-management-api/internal/services/post/delivery"
	postRepository "github.com/jevvonn/post-management-api/internal/services/post/repository"
	postUsecase "github.com/jevvonn/post-management-api/internal/services/post/usecase"
)

func NewHTTPServer() {
	environment := os.Getenv("APP_ENV")
	ports := os.Getenv("APP_PORT")

	router := gin.Default()
	database.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/docs/index.html")
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Repository Initialization
	postRepo := postRepository.NewPostRepository(database.DB)

	// Post Usecase Initialization
	postUsecase := postUsecase.NewPostUsecase(postRepo)

	// Delivery Initialization
	postDelivery.NewPostDelivery(router, postUsecase)

	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		router.Run(fmt.Sprintf(":%s", ports))
	} else {
		router.Run(fmt.Sprintf("localhost:%s", ports))
	}
}
