package main

import (
	"github.com/galang-dana/config"
	"github.com/galang-dana/domain/input"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/domain/usecase"
)

func main() {

	db, _ := config.ConnectDB()
	repo := repository.NewUserRepository(db)

	user := usecase.NewUserUsecase(repo)

	inputsas := input.RegisterUserInput{
		Name:       "jonisca",
		Occupation: "tet",
		Email:      "iniemail@com",
	}
	user.Register(inputsas)
}
