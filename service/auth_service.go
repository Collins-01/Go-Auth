package service

import (
	"errors"
	"net/http"

	"github.com/Collins-01/Go-Auth/dtos"
	"github.com/Collins-01/Go-Auth/models"
	"github.com/Collins-01/Go-Auth/repository"
	"github.com/Collins-01/Go-Auth/utils"
)

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

func (s *AuthService) Login(input *dtos.LoginDTO) (map[string]interface{}, utils.CustomResponse) {
	//Check if a user with this email already exists
	result := s.userRepository.GetUserByEmail(input.Email)
	if result == nil {
		return nil, utils.CustomResponse{
			Code: http.StatusNotFound,
			Err:  errors.New("a user with this email does not exist"),
		}
	}
	// TODO: compare passwords

	// TODO: Generate token with JWT

	return nil, utils.CustomResponse{
		Code: http.StatusOK,
		Err:  nil,
	}

}

func (s *AuthService) SignUp(input *dtos.CreateUserDTO) utils.CustomResponse {
	//Check if a user with the same name already exists
	result := s.userRepository.GetUserByEmail(input.Email)
	if result != nil {
		return utils.CustomResponse{
			Code: http.StatusConflict,
			Err:  errors.New("a user with this email already exists"),
		}
	}
	//Create an instance of the user model
	user := models.UserModel{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		FullName: input.FullName,
	}
	//Perform the insert operation into the user model
	_, status := s.userRepository.CreateUser(&user)
	if status.Err != nil {
		return utils.CustomResponse{
			Code: status.Code,
			Err:  status.Err,
		}
	}
	return utils.CustomResponse{
		Code: http.StatusCreated,
		Err:  nil,
	}

}
