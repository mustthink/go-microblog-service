package service

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/mustthink/microblog/grpc/gen/microblog"
	"github.com/mustthink/microblog/internal/models"
)

type (
	PostStorage interface {
		CreatePost(post *models.Post) error
		GetPostByID(id uint64) (*models.Post, error)
		GetTopicPosts(topicID uint64) ([]models.Post, error)
		DeletePost(id uint64) error
		UpdatePost(post *models.Post) error
	}

	TopicStorage interface {
		CreateTopic(topic *models.Topic) error
		GetTopicByID(id uint64) (*models.Topic, error)
		GetTopicsByAuthor(authorID uint64) ([]models.Topic, error)
		DeleteTopic(id uint64) error
		UpdateTopic(topic *models.Topic) error
	}

	CommentsStorage interface {
		CreateComment(comment *models.Comment) error
		GetCommentByID(id uint64) (*models.Comment, error)
		GetPostComments(postID uint64) ([]models.Comment, error)
		DeleteComment(id uint64) error
		UpdateComment(comment *models.Comment) error
	}

	Microblog struct {
		microblog.UnimplementedMicroblogServer
		postStorage     PostStorage
		topicStorage    TopicStorage
		commentsStorage CommentsStorage
		logger          *logrus.Entry
	}
)

func NewMicroblog(
	postStorage PostStorage,
	topicStorage TopicStorage,
	commentsStorage CommentsStorage,
	logger *logrus.Entry,
) *Microblog {
	return &Microblog{
		postStorage:     postStorage,
		topicStorage:    topicStorage,
		commentsStorage: commentsStorage,
		logger:          logger,
	}
}

func (m *Microblog) Register(server *grpc.Server) {
	microblog.RegisterMicroblogServer(server, m)
}
