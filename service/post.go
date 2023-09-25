package service

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"trail_backend/api/dto"
	"trail_backend/config"
	"trail_backend/models"
	"trail_backend/repository"
	"trail_backend/utils"
)

type (
	PostService interface {
		Create(ctx context.Context, req dto.CreatePostRequest) (*models.Post, error)
		Update(ctx context.Context, req dto.UpdatePostRequest) (*models.Post, error)
		Delete(ctx context.Context, req dto.DeletePostRequest) error
		GetOne(ctx context.Context, req dto.GetOnePostRequest) (*models.Post, error)
		GetList(ctx context.Context, req dto.GetListPostRequest) ([]*models.Post, *int64, error)
	}
	postService struct {
		postRepository repository.PostRepository
		config         *config.Config
	}
)

func NewPostService(itemRepo repository.PostRepository, config *config.Config) PostService {
	return &postService{
		postRepository: itemRepo,
		config:         config,
	}
}

func (s *postService) Create(ctx context.Context, req dto.CreatePostRequest) (*models.Post, error) {
	post := &models.Post{}
	var err error

	if err = utils.Copy(post, req); err != nil {
		return nil, errors.WithStack(err)
	}
	post.UpdaterID = uuid.FromStringOrNil(req.UserId)

	if err = s.postRepository.Create(ctx, post); err != nil {
		return nil, err
	}

	return post, err
}

func (s *postService) Update(ctx context.Context, req dto.UpdatePostRequest) (*models.Post, error) {
	post := &models.Post{}
	var err error

	if err = utils.Copy(post, req); err != nil {
		return nil, err
	}
	post.UpdaterID = uuid.FromStringOrNil(req.UserId)

	if err = s.postRepository.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, err
}

func (s *postService) Delete(ctx context.Context, req dto.DeletePostRequest) error {
	err := s.postRepository.DeleteById(ctx, req)
	return err
}

func (s *postService) GetOne(ctx context.Context, req dto.GetOnePostRequest) (*models.Post, error) {
	post := &models.Post{}
	var err error

	post, err = s.postRepository.GetOneById(ctx, req)
	return post, err
}

func (s *postService) GetList(ctx context.Context, req dto.GetListPostRequest) ([]*models.Post, *int64, error) {
	return s.postRepository.GetList(ctx, req)
}
