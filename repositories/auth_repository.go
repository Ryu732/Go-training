package repositories

import (
	"Gin/models"
	"errors"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) error {
	resulut := r.db.Create(&user)
	if resulut.Error != nil {
		return resulut.Error
	}
	return nil
}

func (r *AuthRepository) FindUser(email string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "email=?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}
