package repository

import (
	"github.com/galang-dana/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(user model.User) (model.User, error)
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
