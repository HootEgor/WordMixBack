package UserService

import (
	Models "WordMixBack/src/Model"
	"context"
)

type UserRepository interface {
	GetUserInfo(ctx context.Context, id string) (Models.User, error)
	GetLeaders(ctx context.Context) ([]Models.Score, error)
	GetUserHistory(ctx context.Context, id string) ([]Models.Score, error)
	NewUserScore(ctx context.Context, score Models.Score) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserInfo(ctx context.Context, id string) (Models.User, error) {
	return s.userRepository.GetUserInfo(ctx, id)
}

func (s *UserService) GetLeaders(ctx context.Context) ([]Models.Score, error) {
	return s.userRepository.GetLeaders(ctx)
}

func (s *UserService) GetUserHistory(ctx context.Context, id string) ([]Models.Score, error) {
	return s.userRepository.GetUserHistory(ctx, id)
}

func (s *UserService) NewUserScore(ctx context.Context, score Models.Score) error {
	return s.userRepository.NewUserScore(ctx, score)
}
