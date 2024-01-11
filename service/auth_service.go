package service

import "github.com/Collins-01/Go-Auth/repository"

// Authentication service
type AuthService struct {
	userRepository repository.UserRepository
}

// Returns and instance of AuthService
func NewAuthService(userRepository repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Login() {}

func (s *AuthService) SignUp() {}
