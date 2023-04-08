package main

import (
	"github.com/galang-dana/database"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/interfaces/delivery/auth"
	"github.com/galang-dana/interfaces/delivery/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := database.Connect()
	repo := repository.NewUserRepository(db)

	user := usecase.NewUserUsecase(repo)
	auth := auth.NewService()

	// login := input.CheckEmail{
	// 	Email: "rachel@mail.com",
	// }
	// usr, err := user.EmailAvaliable(login)
	// if err != nil {
	// 	fmt.Println("ada error")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(usr)
	// fmt.Println("Bisa digunakan")

	userHandler := handler.NewUserHandler(user, auth)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email-check", userHandler.CheckEmail)
	api.POST("/upload-avatar", userHandler.UploadAvatar)

	router.Run()

}
