package main

import (
	"fmt"

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

	cek, _ := repoCampaign.FindById("f243c33d14244256bb122a5015834e15")
	for _, v := range cek {

		fmt.Println(v.Name)
		if len(v.Images) > 0 {
			fmt.Println(v.Images[0].FileName)
			fmt.Println("Jumlah gambar : ", len(v.Images))
		}
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
