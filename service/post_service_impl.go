package service

import (
	"context"
	"go_prisma/data/request"
	"go_prisma/data/response"
	"go_prisma/helper"
	"go_prisma/model"
	"go_prisma/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepository repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository}
}

func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	postData := model.Post{
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	postData := model.Post{
		Id:          request.Id,
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Update(ctx, postData)
}

func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)
}

func (p *PostServiceImpl) FindByID(ctx context.Context, postId string) response.PostResponse {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)

	postResponse := response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}
	return postResponse
}

func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	posts := p.PostRepository.FindAll(ctx)

	var postsResponse []response.PostResponse

	for _, value := range posts {
		post := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Published:   value.Published,
			Description: value.Description,
		}
		postsResponse = append(postsResponse, post)
	}
	return postsResponse
}
