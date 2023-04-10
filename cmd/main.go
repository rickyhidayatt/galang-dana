package main

import (
	"fmt"

	"github.com/galang-dana/database"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/domain/usecase"
)

func main() {

	db, _ := database.Connect()
	// repo := repository.NewUserRepository(db)

	// user := usecase.NewUserUsecase(repo)
	// auth := auth.NewService()

	//camapaign
	repoCampaign := repository.NewCampaignRepository(db)
	usecaseCampaign := usecase.NewCampaignUseCase(repoCampaign)

	cek, _ := usecaseCampaign.FindCampaigns("c7a626ab48cc419399b5c662ea6a9043")
	for _, v := range cek {

		fmt.Println("Nama camapaign : ", v.Name)

	}

	// userHandler := handler.NewUserHandler(user, auth)

	// router := gin.Default()
	// api := router.Group("api/v1")

	// api.POST("/users", userHandler.RegisterUser)
	// api.POST("/login", userHandler.LoginUser)
	// api.POST("/email-check", userHandler.CheckEmail)
	// api.POST("/upload-avatar", middleware.AuthMiddleware(auth, user), userHandler.UploadAvatar)

	// router.Run()

}
