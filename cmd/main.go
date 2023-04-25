package main

import (
	"github.com/galang-dana/database"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/interfaces/delivery/auth"
	"github.com/galang-dana/interfaces/delivery/handler"
	"github.com/galang-dana/interfaces/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := database.Connect()
	repo := repository.NewUserRepository(db)

	user := usecase.NewUserUsecase(repo)
	auth := auth.NewService()
	userHandler := handler.NewUserHandler(user, auth)

	//camapaign
	repoCampaign := repository.NewCampaignRepository(db)
	CampaignUsecase := usecase.NewCampaignUseCase(repoCampaign)
	campaignHandler := handler.CampaignHandler(CampaignUsecase)

	router := gin.Default()
	router.Static("/images", "../images")
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email-check", userHandler.CheckEmail)
	api.POST("/upload-avatar", middleware.AuthMiddleware(auth, user), userHandler.UploadAvatar)

	// campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(auth, user), campaignHandler.CreateCampaign)

	api.GET("/")

	router.Run()

}
