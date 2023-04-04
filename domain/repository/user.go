package repository

import (
	"github.com/galang-dana/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(user model.User) (model.User, error)
	FindEmail(email string) (model.User, error)
	FindByID(id string) (model.User, error)
	CheckEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(g *gorm.DB) UserRepository {
	return &userRepository{g}
}

// Insert
func (u *userRepository) SaveUser(user model.User) (model.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) FindEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) FindByID(id string) (model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) CheckEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
