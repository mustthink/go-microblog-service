package storage

import (
	"fmt"

	"github.com/mustthink/microblog/internal/models"
)

func (s *Storage) CreateComment(comment *models.Comment) error {
	const op = "storage.CreateComment"

	if err := s.db.Create(comment).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetCommentByID(id uint64) (*models.Comment, error) {
	const op = "storage.GetCommentByID"

	comment := new(models.Comment)
	if err := s.db.First(comment, id).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return comment, nil
}

func (s *Storage) GetPostComments(postID uint64) ([]models.Comment, error) {
	const op = "storage.GetPostComments"

	var comments []models.Comment
	if err := s.db.Find(&comments, "post_id = ?", postID).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return comments, nil
}

func (s *Storage) DeleteComment(id uint64) error {
	const op = "storage.DeleteComment"

	if err := s.db.Delete(&models.Comment{}, id).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) UpdateComment(comment *models.Comment) error {
	const op = "storage.UpdateComment"

	if err := s.db.Save(comment).Error; err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
