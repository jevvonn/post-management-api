package post

import (
	"github.com/jevvonn/post-management-api/internal/model/domain"
	"github.com/jevvonn/post-management-api/internal/model/dto"
)

type PostUsecase interface {
	GetPosts() ([]domain.Post, error)
	CreatePost(req *dto.CreatePostRequest) error
	GetPostByID(id int) (domain.Post, error)
}
