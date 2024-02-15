package storage

import (
	"fmt"

	"github.com/mustthink/microblog/internal/models"
)

func (s *Storage) CreateTopic(topic *models.Topic) error {
	const op = "storage.CreateTopic"

	if err := s.db.Create(topic).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetTopicByID(id uint64) (*models.Topic, error) {
	const op = "storage.GetTopicByID"

	var topic models.Topic
	if err := s.db.First(&topic, id).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &topic, nil
}

func (s *Storage) GetTopicsByAuthor(authorID uint64) ([]models.Topic, error) {
	const op = "storage.GetTopicsByAuthor"

	var topics []models.Topic
	if err := s.db.Find(&topics, "author_id = ?", authorID).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return topics, nil
}

func (s *Storage) DeleteTopic(id uint64) error {
	const op = "storage.DeleteTopic"

	if err := s.db.Delete(&models.Topic{}, id).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) UpdateTopic(topic *models.Topic) error {
	const op = "storage.UpdateTopic"

	if err := s.db.Save(topic).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
