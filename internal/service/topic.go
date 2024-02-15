package service

import (
	"context"

	"github.com/mustthink/microblog/grpc/gen/microblog"
	"github.com/mustthink/microblog/internal/models"
)

func (m *Microblog) CreateTopic(_ context.Context, req *microblog.CreateTopicRequest) (*microblog.CreateTopicResponse, error) {
	op := "service.CreateTopic"
	log := m.logger.WithField("op", op)
	topic := &models.Topic{
		Title:   req.GetTitle(),
		Content: req.GetDescription(),
	}

	if err := m.topicStorage.CreateTopic(topic); err != nil {
		log.Debugf("failed to create topic: %s", err.Error())
		return &microblog.CreateTopicResponse{TopicId: 0}, err
	}

	return &microblog.CreateTopicResponse{TopicId: uint64(topic.ID)}, nil
}

func (m *Microblog) GetTopic(_ context.Context, req *microblog.GetTopicRequest) (*microblog.GetTopicResponse, error) {
	op := "service.GetTopic"
	log := m.logger.WithField("op", op)

	topic, err := m.topicStorage.GetTopicByID(req.GetTopicId())
	if err != nil {
		log.Debugf("failed to get topic: %s", err.Error())
		return nil, err
	}

	return &microblog.GetTopicResponse{
		Title:       topic.Title,
		Description: topic.Content,
	}, nil
}

func (m *Microblog) GetTopicPosts(_ context.Context, req *microblog.TopicPostsRequest) (*microblog.TopicPostsResponse, error) {
	op := "service.GetTopicPosts"
	log := m.logger.WithField("op", op)

	posts, err := m.postStorage.GetTopicPosts(req.GetTopicId())
	if err != nil {
		log.Debugf("failed to get topic posts: %s", err.Error())
		return nil, err
	}

	var resp = make([]*microblog.GetPostResponse, 0, len(posts))
	for _, post := range posts {
		resp = append(resp, &microblog.GetPostResponse{
			TopicId:  post.TopicID,
			Title:    post.Title,
			Content:  post.Content,
			AuthorId: post.AuthorID,
		})
	}

	return &microblog.TopicPostsResponse{Posts: resp}, nil
}

func (m *Microblog) DeleteTopic(_ context.Context, req *microblog.DeleteTopicRequest) (*microblog.DeleteTopicResponse, error) {
	op := "service.DeleteTopic"
	log := m.logger.WithField("op", op)

	if err := m.topicStorage.DeleteTopic(req.GetTopicId()); err != nil {
		log.Debugf("failed to delete topic: %s", err.Error())
		return &microblog.DeleteTopicResponse{}, err
	}

	return &microblog.DeleteTopicResponse{
		Success: true,
	}, nil
}
