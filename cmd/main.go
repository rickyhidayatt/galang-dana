package main

import (
	"github.com/galang-dana/database"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/interfaces/delivery/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := database.Connect()
	repo := repository.NewUserRepository(db)

	user := usecase.NewUserUsecase(repo)

	userHandler := handler.NewUserHandler(user)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run()

}
