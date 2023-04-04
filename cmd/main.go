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

	// login := input.LoginUser{
	// 	Email:    "rachesl@mail.com",
	// 	Password: "rachel",
	// }
	// usr, err := user.Login(login)
	// if err != nil {
	// 	fmt.Println("ada error")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(usr.Email)
	// fmt.Println(usr.Name)

	userHandler := handler.NewUserHandler(user)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)

	router.Run()

}
