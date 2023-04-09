package main

import (
	"github.com/galang-dana/database"
	"github.com/galang-dana/domain/repository"
)

func main() {

	db, _ := database.Connect()
	// repo := repository.NewUserRepository(db)

	// user := usecase.NewUserUsecase(repo)
	// auth := auth.NewService()

	//camapaign
	repoCampaign := repository.NewCampaignRepository(db)

	repoCampaign.FindAll()

	// userHandler := handler.NewUserHandler(user, auth)

	// router := gin.Default()
	// api := router.Group("api/v1")

	// api.POST("/users", userHandler.RegisterUser)
	// api.POST("/login", userHandler.LoginUser)
	// api.POST("/email-check", userHandler.CheckEmail)
	// api.POST("/upload-avatar", middleware.AuthMiddleware(auth, user), userHandler.UploadAvatar)

	// router.Run()

}
