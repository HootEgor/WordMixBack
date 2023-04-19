package Services

import (
	Models "WordMixBack/src/Model"
	"context"
)

type AuthRepository interface {
	AddNewUser(ctx context.Context, user Models.User) (string, error)
	GetUserByInfo(ctx context.Context, user Models.User) (string, error)
}

type UserRepository interface {
	GetUserInfo(ctx context.Context, id string) (Models.User, error)
	GetLeaders(ctx context.Context) ([]Models.Score, error)
	GetUserHistory(ctx context.Context, id string) ([]Models.Score, error)
	NewUserScore(ctx context.Context, score Models.Score) error
}

type UserService struct {
	authRepository AuthRepository
	userRepository UserRepository
}

func NewUserService(authRepository AuthRepository, userRepository UserRepository) *UserService {
	return &UserService{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

// AuthService
func (s *UserService) AddNewUser(ctx context.Context, user Models.User) (string, error) {
	return s.authRepository.AddNewUser(ctx, user)
}

func (s *UserService) LoginUser(ctx context.Context, user Models.User) (string, error) {
	login, err := s.authRepository.GetUserByInfo(ctx, user)
	if err != nil {
		return "", err
	}
	return login, nil
}

// UserService
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
