package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/jevvonn/post-management-api/internal/services/post"
)

type PostDelivery struct {
	router      *gin.Engine
	repoUsecase post.PostUsecase
}

func NewPostDelivery(router *gin.Engine, repoUsecase post.PostUsecase) {
	handler := &PostDelivery{router, repoUsecase}

	router.GET("/posts", handler.GetPosts)
	// router.POST("/posts", handler.CreatePost)
}

func (p *PostDelivery) GetPosts(c *gin.Context) {
	posts, err := p.repoUsecase.GetPosts()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"errors":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success get posts data!",
		"data":    posts,
	})
}

// func (p *PostDelivery) CreatePost(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Hello World",
// 	})
// }
