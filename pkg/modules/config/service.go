package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// ErrKeyNotExists indicates the env key does not exists
	ErrKeyNotExists = errors.New("env key does not exists")
)

type Service struct {
	config Config
}

func NewService(c Config) (*Service, error) {
	err := godotenv.Load(c.envFilePaths...)
	if err != nil {
		return nil, err
	}

	return &Service{
		config: c,
	}, nil
}

func (s *Service) Get(key string) string { return os.Getenv(key) }

func (s *Service) Lookup(key string) (string, bool) { return os.LookupEnv(key) }

func (s *Service) MustGet(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("%w: %v does not exists", ErrKeyNotExists, key))
	}
	return val
}
