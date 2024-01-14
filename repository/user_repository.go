package repository

import (
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
	db := r.db
	result := db.Create(&user)
	if result.Error != nil {
		return nil, utils.CustomResponse{
			Code: http.StatusInternalServerError,
		}
	}
	return user, utils.CustomResponse{
		Err:  nil,
		Code: http.StatusCreated,
	}
}

func (r *UserRepository) GetUserByEmail(email string) *models.UserModel {
	db := r.db
	var user models.UserModel
	err := db.First(&user).Where("email = ?", email).Error
	if err != nil {
		return nil
	}
	return &user
}

func (r *UserRepository) GetUserByID(id string) *models.UserModel {
	db := r.db
	var user models.UserModel
	err := db.First(&user).Where("id = ?", id).Error
	if err != nil {
		return nil
	}
	return &user
}
