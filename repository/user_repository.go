package repository

import (
	"errors"
	"net/http"

	"github.com/Collins-01/Go-Auth/models"
	"github.com/Collins-01/Go-Auth/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	//Inject gorm library into the repository
	db *gorm.DB
}

// Method to return instance of the user repository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// A method to create a new user on the database. Returns `409` if a user with the same email already exists.
func (r *UserRepository) CreateUser(user *models.UserModel) (*models.UserModel, utils.CustomResponse) {

	return nil, utils.CustomResponse{
		Err:  errors.New("a user with this email already exists"),
		Code: http.StatusConflict,
	}
}

func (r *UserRepository) GetUserByEmail(email string) *models.UserModel {
	return nil
}

func (r *UserRepository) GetUserByID(id string) *models.UserModel {
	return nil
}
