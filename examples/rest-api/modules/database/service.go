package database

import (
	"fmt"

	"github.com/huboh/vara/pkg/modules/config"
)

type Service struct {
	configs *config.Service
}

func newService(c *config.Service) (*Service, error) {
	db := &Service{configs: c}
	err := db.connect(c.MustGet("DB_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func (db *Service) connect(string) error {
	return nil
}

func (db *Service) Transaction(f func() error) error {
	err := f()
	if err != nil {
		fmt.Errorf("database transaction failed: %w", err)
	}

	return nil
}
