package config

import (
	"os"
	"path"
)

type Config struct {
	cache        bool
	envFilePaths []string
}

func NewConfig() (Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	return Config{
		cache:        false,
		envFilePaths: []string{path.Join(wd, ".env")},
	}, nil
}
