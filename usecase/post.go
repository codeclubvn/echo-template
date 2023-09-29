package usecase

import (
	"context"
	"echo_template/config"
	"echo_template/domain/repo"
	"echo_template/domain/repo/model"
	"echo_template/pkg/api_errors"
	"echo_template/pkg/utils"
	"echo_template/presenter/request"
	"github.com/pkg/errors"
)

type (
	PostService interface {
		Create(ctx context.Context, req request.CreatePostRequest) (*model.Post, error)
		Update(ctx context.Context, req request.UpdatePostRequest) (*model.Post, error)
		Delete(ctx context.Context, req request.DeletePostRequest) error
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
	post, err := s.GetOne(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	// check post is belong to user
	if post.UserId != req.UserId {
		return nil, errors.New(api_errors.ErrUnauthorizedAccess)
	}

	if err := utils.Copy(post, req); err != nil {
		return nil, err
	}
	if err := s.postRepository.Update(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) Delete(ctx context.Context, req request.DeletePostRequest) error {
	post, err := s.GetOne(ctx, req.ID)
	if err != nil {
		return err
	}
	// check post is belong to user
	if post.UserId != req.UserId {
		return errors.New(api_errors.ErrUnauthorizedAccess)
	}
	return s.postRepository.DeleteById(ctx, req.ID)
}

func (s *postService) GetOne(ctx context.Context, id string) (*model.Post, error) {
	return s.postRepository.GetOneById(ctx, id)
}

func (s *postService) GetList(ctx context.Context, req request.GetListPostRequest) ([]*model.Post, *int64, error) {
	return s.postRepository.GetList(ctx, req)
}
