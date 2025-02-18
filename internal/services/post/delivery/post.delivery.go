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
	apiRouter := router.Group("/api")

	apiRouter.GET("/api/posts", handler.GetPosts)
	apiRouter.GET("/api/posts/:id", handler.GetPostByID)
}

// @title 			Get All Posts
//
//	@Tags			Posts
//	@Summary		Get All Posts
//	@Description	Get All Posts
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]domain.Post{}
//
// @Security BearerAuth
//
//	@Router			/api/posts [get]
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

// @title 			Get Post By Id
//
//	@Tags			Posts
//	@Summary		Get Post By Id
//	@Description	Get Post By Id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	domain.Post{}
//
// @Param			id		path	int	true	"Post ID"
//
// @Security BearerAuth
//
//	@Router			/api/posts/{id} [get]
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
