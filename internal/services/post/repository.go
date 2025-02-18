package post

import "github.com/jevvonn/post-management-api/internal/model/domain"

type PostRepository interface {
	GetPosts() ([]domain.Post, error)
	CreatePost(post domain.Post) error
	GetPostByID(id int) (domain.Post, error)
	UpdatePost(post domain.Post) error
	DeletePost(id int) error
}
