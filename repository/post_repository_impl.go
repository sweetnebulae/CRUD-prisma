package repository

import (
	"context"
	"fmt"
	"go_prisma/helper"
	"go_prisma/model"
	"go_prisma/prisma/db"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

func NewPostRepositoryImpl(Db *db.PrismaClient) *PostRepositoryImpl {
	return &PostRepositoryImpl{Db: Db}
}

func (p *PostRepositoryImpl) Delete(ctx context.Context, postID string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postID)).Delete().Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Row affected: ", result)
}

func (p *PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {

	allPosts, err := p.Db.Post.FindMany().Exec(ctx)
	helper.ErrorPanic(err)

	var posts []model.Post

	for _, post := range allPosts {
		published, _ := post.Published()
		description, _ := post.Description()

		postData := model.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: description,
		}
		posts = append(posts, postData)
	}
	return posts
}

func (p *PostRepositoryImpl) FindById(ctx context.Context, postID string) ([]model.Post, error) {
	post, err := p.Db.Post.FindFirst(db.Post.ID.Equals(postID)).Exec(ctx)
	helper.ErrorPanic(err)

	published, _ := post.Published()
	description, _ := post.Description()
	postData := model.Post{
		Id:          post.ID,
		Title:       post.Title,
		Published:   published,
		Description: description,
	}

	if post != nil {
		return postData, nil
	} else {
		return postData, error.New("Post id not found")
	}
}

func (*PostRepositoryImpl) FindByUserId(ctx context.Context, userID string) ([]model.Post, error) {}
