package AuthService

import (
	Models "WordMixBack/src/Model"
	"context"
)

type AuthRepository interface {
	AddNewUser(ctx context.Context, user Models.User) (string, error)
	GetUserByInfo(ctx context.Context, user Models.User) (string, error)
}

type AuthService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (s *AuthService) AddNewUser(ctx context.Context, user Models.User) (string, error) {
	return s.authRepository.AddNewUser(ctx, user)
}

func (s *AuthService) LoginUser(ctx context.Context, user Models.User) (string, error) {
	login, err := s.authRepository.GetUserByInfo(ctx, user)
	if err != nil {
		return "", err
	}
	return login, nil
}
