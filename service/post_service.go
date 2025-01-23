package service

import (
	"context"
	"go_prisma/data/request"
	"go_prisma/data/response"
)

type PostService interface {
	Create(ctx context.Context, request request.PostCreateRequest)
	Update(ctx context.Context, request request.PostUpdateRequest)
	Delete(ctx context.Context, postId string)
	FindByID(ctx context.Context, postId string) response.PostResponse
	FindAll(ctx context.Context) []response.PostResponse
}
