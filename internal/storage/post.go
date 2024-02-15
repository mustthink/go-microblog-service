package storage

import (
	"fmt"

	"github.com/mustthink/microblog/internal/models"
)

func (s *Storage) CreatePost(post *models.Post) error {
	const op = "storage.CreatePost"

	if err := s.db.Create(post).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetPostByID(id uint64) (*models.Post, error) {
	const op = "storage.GetPostByID"

	post := new(models.Post)
	if err := s.db.First(post, id).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return post, nil
}

func (s *Storage) GetTopicPosts(topicID uint64) ([]models.Post, error) {
	const op = "storage.GetTopicPosts"

	var posts []models.Post
	if err := s.db.Find(&posts, "topic_id = ?", topicID).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (s *Storage) DeletePost(id uint64) error {
	const op = "storage.DeletePost"

	if err := s.db.Delete(&models.Post{}, id).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) UpdatePost(post *models.Post) error {
	const op = "storage.UpdatePost"

	if err := s.db.Save(post).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
