package usecase

import (
	"errors"

	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/model"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(input input.RegisterUserInput) (model.User, error)
	Login(input input.LoginUser) (model.User, error)
	EmailAvaliable(input input.CheckEmail) (bool, error)
	SaveAvatar(Id string, FileLocation string) (model.User, error)
	GetUserById(id string) (model.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(u repository.UserRepository) UserUseCase {
	return &userUseCase{u}
}

func (s *userUseCase) Register(input input.RegisterUserInput) (model.User, error) {

	var user = model.User{}
	user.Id = utils.Uuid()
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.userRepo.SaveUser(user)
	if err != nil {
		return newUser, errors.New("faild register user")
	}

	return newUser, nil
}

func (s *userUseCase) Login(input input.LoginUser) (model.User, error) {
	email := input.Email
	pasword := input.Password

	user, err := s.userRepo.FindEmail(email)
	if err != nil {
		return user, errors.New("no user found on that email")
	}

	checkId, err := s.userRepo.FindByID(user.Id)
	if err != nil {
		return checkId, errors.New("no user id found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pasword))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userUseCase) EmailAvaliable(input input.CheckEmail) (bool, error) {
	email := input.Email
	user, err := s.userRepo.CheckEmail(email)
	if err != nil {
		return false, errors.New("e-mail has been used")
	}

	if user.Id == "" {
		return true, nil
	}

	return false, nil
}

func (s *userUseCase) SaveAvatar(Id string, FileLocation string) (model.User, error) {

	user, err := s.userRepo.FindByID(Id)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = FileLocation
	userUpdate, err := s.userRepo.Update(user)
	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *userUseCase) GetUserById(id string) (model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return user, err
	}

	if user.Id == "" {
		return user, errors.New("no user found with that id")
	}

	return user, nil
}
