package user

import (
	"context"
	"rest-api-project/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	// TODO next
	return User{}, nil
}
