package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jevvonn/post-management-api/database"
	postDelivery "github.com/jevvonn/post-management-api/internal/services/post/delivery"
)

func NewHTTPServer() {
	environment := os.Getenv("APP_ENV")
	ports := os.Getenv("APP_PORT")

	router := gin.Default()
	database.ConnectDatabase()

	// Delivery Initialization
	postDelivery.NewPostDelivery(router)

	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		router.Run(fmt.Sprintf(":%s", ports))
	} else {
		router.Run(fmt.Sprintf("localhost:%s", ports))
	}
}
