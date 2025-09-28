package service

import (
	"github.com/prankevich/MyProject/internal/contracts"
	"github.com/rs/zerolog"
)

type Service struct {
	repository contracts.RepositoryI
	cache      contracts.CacheI
	logger     zerolog.Logger
}

func NewService(repository contracts.RepositoryI, cache contracts.CacheI, logger zerolog.Logger) *Service {
	return &Service{
		repository: repository,
		cache:      cache,
		logger:     logger,
	}
}
