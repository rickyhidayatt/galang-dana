package delivery

import (
	"github.com/galang-dana/database"
	"github.com/galang-dana/domain/repository"
	"github.com/galang-dana/domain/usecase"
	"github.com/galang-dana/interfaces/delivery/auth"
	"github.com/galang-dana/interfaces/delivery/handler"
	"github.com/galang-dana/interfaces/delivery/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {

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

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "../images")
	api := router.Group("api/v1")

	// user Endpoint
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email-check", userHandler.CheckEmail)
	api.POST("/upload-avatar", middleware.AuthMiddleware(auth, user), userHandler.UploadAvatar)
	api.GET("/users/fetch", middleware.AuthMiddleware(auth, user), userHandler.FetchUser)

	// campaign Endpoint
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignById)
	api.POST("/campaigns", middleware.AuthMiddleware(auth, user), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(auth, user), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(auth, user), campaignHandler.UploadImage)

	// transaction Endpoint
	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(auth, user), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", middleware.AuthMiddleware(auth, user), transactionHandler.GetUserTransactions)
	api.POST("/transactions", middleware.AuthMiddleware(auth, user), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	api.GET("/")
	router.Run()
}
