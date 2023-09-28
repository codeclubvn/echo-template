package usecase

import (
	"context"
	"github.com/pkg/errors"
	"trial_backend/config"
	"trial_backend/domain/repo"
	"trial_backend/domain/repo/model"
	"trial_backend/pkg/utils"
	"trial_backend/presenter/request"
)

type (
	PostService interface {
		Create(ctx context.Context, req request.CreatePostRequest) (*model.Post, error)
		Update(ctx context.Context, req request.UpdatePostRequest) (*model.Post, error)
		Delete(ctx context.Context, id string) error
		GetOne(ctx context.Context, id string) (*model.Post, error)
		GetList(ctx context.Context, req request.GetListPostRequest) ([]*model.Post, *int64, error)
	}
	postService struct {
		postRepository repo.PostRepository
		config         *config.Config
	}
)

func NewPostService(itemRepo repo.PostRepository, config *config.Config) PostService {
	return &postService{
		postRepository: itemRepo,
		config:         config,
	}
}

func (s *postService) Create(ctx context.Context, req request.CreatePostRequest) (*model.Post, error) {
	post := &model.Post{}
	if err := utils.Copy(post, req); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := s.postRepository.Create(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) Update(ctx context.Context, req request.UpdatePostRequest) (*model.Post, error) {
	post := &model.Post{}
	if err := utils.Copy(post, req); err != nil {
		return nil, err
	}
	if err := s.postRepository.Update(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) Delete(ctx context.Context, id string) error {
	return s.postRepository.DeleteById(ctx, id)
}

func (s *postService) GetOne(ctx context.Context, id string) (*model.Post, error) {
	return s.postRepository.GetOneById(ctx, id)
}

func (s *postService) GetList(ctx context.Context, req request.GetListPostRequest) ([]*model.Post, *int64, error) {
	return s.postRepository.GetList(ctx, req)
}
