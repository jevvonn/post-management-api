package usecase

import (
	"github.com/jevvonn/post-management-api/helper"
	"github.com/jevvonn/post-management-api/internal/model/domain"
	"github.com/jevvonn/post-management-api/internal/model/dto"
	"github.com/jevvonn/post-management-api/internal/services/post"
)

type PostUsecase struct {
	postRepo post.PostRepository
}

func NewPostUsecase(postRepo post.PostRepository) post.PostUsecase {
	return &PostUsecase{postRepo: postRepo}
}

func (p *PostUsecase) GetPosts() ([]domain.Post, error) {
	posts, err := p.postRepo.GetPosts()
	if err != nil {
		return nil, nil
	}

	return posts, nil
}

func (p *PostUsecase) CreatePost(req *dto.CreatePostRequest) error {
	dateBeforeRevision, err := helper.StringISOToDateTime(req.DeadlineBeforeRevision)
	dateAfterRevision, err := helper.StringISOToDateTime(req.DeadlineAfterRevision)
	uploadDate, err := helper.StringISOToDateTime(req.UploadDate)

	if err != nil {
		return err
	}

	post := domain.Post{
		Title:                  req.Title,
		ContentDescription:     req.ContentDescription,
		Caption:                req.Caption,
		Platforms:              req.Platforms,
		DesignLink:             req.DesignLink,
		DeadlineBeforeRevision: dateBeforeRevision,
		DeadlineAfterRevision:  dateAfterRevision,
		UploadDate:             uploadDate,
		Status:                 req.Status,
	}

	err = p.postRepo.CreatePost(post)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostUsecase) GetPostByID(id int) (domain.Post, error) {
	post, err := p.postRepo.GetPostByID(id)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}
