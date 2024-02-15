package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/mustthink/microblog/internal/config"
)

type Storage struct {
	db *gorm.DB
}

func generateURI(cfg config.DB) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password)
}

func New(cfg config.DB) (*Storage, error) {
	const op = "storage.New"

	db, err := gorm.Open("postgres", generateURI(cfg))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
