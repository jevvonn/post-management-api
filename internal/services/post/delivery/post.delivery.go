package delivery

import (
	"strconv"

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
	router.GET("/posts/:id", handler.GetPostByID)
	// router.POST("/posts", handler.CreatePost)
}

func (p *PostDelivery) GetPosts(c *gin.Context) {
	posts, err := p.repoUsecase.GetPosts()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"errors":  err.Error(),
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

func (p *PostDelivery) GetPostByID(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
			"errors":  err.Error(),
		})
		return
	}

	post, err := p.repoUsecase.GetPostByID(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"errors":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success get post data!",
		"data":    post,
	})
}
