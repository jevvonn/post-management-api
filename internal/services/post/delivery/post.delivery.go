package delivery

import (
	"github.com/gin-gonic/gin"
)

type PostDelivery struct {
	router *gin.Engine
}

func NewPostDelivery(router *gin.Engine) {
	handler := &PostDelivery{router: router}

	router.GET("/posts", handler.GetPosts)
	router.POST("/posts", handler.CreatePost)
}

func (p *PostDelivery) GetPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (p *PostDelivery) CreatePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
