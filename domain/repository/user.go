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
	Update(input model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(g *gorm.DB) UserRepository {
	return &userRepository{g}
}

// Insert
func (r *userRepository) SaveUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id string) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) CheckEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
