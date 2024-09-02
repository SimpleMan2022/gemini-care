package repository

import (
	"gemini-care/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(User *entity.User) error
	FindOneByUsername(username string) (*entity.User, error)
	FindOneByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(User *entity.User) error {
	return r.db.Create(User).Error
}

func (r *userRepository) FindOneByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindOneByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
