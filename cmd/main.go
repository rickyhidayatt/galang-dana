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

	//USER
	user := usecase.NewUserUsecase(repo)
	auth := auth.NewService()
	userHandler := handler.NewUserHandler(user, auth)

	//CAMPAIGN
	repoCampaign := repository.NewCampaignRepository(db)
	CampaignUsecase := usecase.NewCampaignUseCase(repoCampaign)
	campaignHandler := handler.CampaignHandler(CampaignUsecase)

	//TRANSACTION
	repoTransaction := repository.NewTransactionRepository(db)

	paymentUseCase := usecase.NewPaymentUseCase(repoTransaction, repoCampaign)
	transactionUseCase := usecase.NewTransactionUseCase(repoTransaction, repoCampaign, paymentUseCase)
	transactionHandler := handler.TransactionHandler(transactionUseCase, paymentUseCase)

	// uid, _ := user.GetUserById("b20457abec1c4773a3ac3010a1114992")
	// input := input.CreateTransactionInput{
	// 	CampaignID: "tse12",
	// 	Amount:     200000,
	// 	User:       uid,
	// }

	// _, err := transactionUseCase.CreateTransaction(input)
	// if err != nil {
	// 	fmt.Println("ERORRRRR CUKKKK")
	// 	fmt.Println("ERORRRRR CUKKKK")
	// 	fmt.Println("ERORRRRR CUKKKK")
	// } else {
	// 	fmt.Println("Berhasil Cok")
	// 	fmt.Println("Berhasil Cok")
	// 	fmt.Println("Berhasil Cok")
	// }

	router := gin.Default()
	router.Static("/images", "../images")
	api := router.Group("api/v1")

	// user Endpoint
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email-check", userHandler.CheckEmail)
	api.POST("/upload-avatar", middleware.AuthMiddleware(auth, user), userHandler.UploadAvatar)

	// campaign Endpoint
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignById)
	api.POST("/campaigns", middleware.AuthMiddleware(auth, user), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(auth, user), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(auth, user), campaignHandler.UploadImage)

	// transaction ENDPOINT

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(auth, user), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", middleware.AuthMiddleware(auth, user), transactionHandler.GetUserTransactions)
	api.POST("/transactions", middleware.AuthMiddleware(auth, user), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	api.GET("/")

	router.Run()

}
