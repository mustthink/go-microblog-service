package service

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/mustthink/microblog/grpc/gen/microblog"
	"github.com/mustthink/microblog/internal/models"
)

func (m *Microblog) CreatePost(_ context.Context, req *microblog.CreatePostRequest) (*microblog.CreatePostResponse, error) {
	op := "service.CreatePost"
	log := m.logger.WithField("op", op)
	post := &models.Post{
		AuthorID: req.GetAuthorId(),
		TopicID:  req.GetTopicId(),
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
	}

	if err := m.postStorage.CreatePost(post); err != nil {
		log.Debugf("failed to create post: %s", err.Error())
		return &microblog.CreatePostResponse{PostId: 0}, err
	}

	return &microblog.CreatePostResponse{PostId: uint64(post.ID)}, nil
}
func (m *Microblog) GetPost(_ context.Context, req *microblog.GetPostRequest) (*microblog.GetPostResponse, error) {
	op := "service.GetPost"
	log := m.logger.WithField("op", op)

	post, err := m.postStorage.GetPostByID(req.GetPostId())
	if err != nil {
		log.Debugf("failed to get post: %s", err.Error())
		return nil, err
	}

	return &microblog.GetPostResponse{
		TopicId:  post.TopicID,
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: post.AuthorID,
	}, nil
}
func (m *Microblog) DeletePost(_ context.Context, req *microblog.DeletePostRequest) (*microblog.DeletePostResponse, error) {
	op := "service.DeletePost"
	log := m.logger.WithField("op", op)

	if err := m.postStorage.DeletePost(req.GetPostId()); err != nil {
		log.Debugf("failed to delete post: %s", err.Error())
		return &microblog.DeletePostResponse{Success: false}, err
	}

	return &microblog.DeletePostResponse{Success: true}, nil
}
func (m *Microblog) UpdatePost(_ context.Context, req *microblog.UpdatePostRequest) (*microblog.UpdatePostResponse, error) {
	op := "service.UpdatePost"
	log := m.logger.WithField("op", op)
	post := &models.Post{
		Model: gorm.Model{
			ID: uint(req.GetPostId()),
		},
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	if err := m.postStorage.UpdatePost(post); err != nil {
		log.Debugf("failed to update post: %s", err.Error())
		return &microblog.UpdatePostResponse{Success: false}, err
	}

	return &microblog.UpdatePostResponse{Success: true}, nil
}
