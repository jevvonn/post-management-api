package repository

import (
	"github.com/jevvonn/post-management-api/internal/model/domain"
	"github.com/jevvonn/post-management-api/internal/services/post"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) post.PostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) GetPosts() ([]domain.Post, error) {
	var posts []domain.Post
	err := p.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostRepository) CreatePost(post domain.Post) error {
	err := p.db.Create(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PostRepository) GetPostByID(id int) (domain.Post, error) {
	var post domain.Post
	err := p.db.Where("id = ?", id).First(&post).Error
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

func (p *PostRepository) UpdatePost(post domain.Post) error {
	err := p.db.Save(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PostRepository) DeletePost(id int) error {
	err := p.db.Delete(&domain.Post{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
