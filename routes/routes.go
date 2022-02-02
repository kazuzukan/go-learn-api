package routes

import (
	"bwa-project/auth"
	"bwa-project/campaign"
	"bwa-project/config"
	"bwa-project/handler"
	"bwa-project/middleware"
	"bwa-project/models"
	"bwa-project/payment"
	"bwa-project/transaction"
	"bwa-project/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	db := config.DbConfig()

	// migrate db
	for _, model := range models.RegisterModel() {
		db.AutoMigrate(model.Model)
	}

	// repository
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// service
	userService := user.NewServices(userRepository)
	authService := auth.NewServices()
	campaignService := campaign.NewServices(campaignRepository)
	paymentService := payment.NewService(campaignRepository)
	transctionService := transaction.NewServices(transactionRepository, campaignRepository, paymentService)

	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHanlder(campaignService)
	transctionHandler := handler.NewtransctionHandler(transctionService)

	router := gin.Default()
	router.Use(cors.Default())
	// param pertama itu yang mau dituju, yang kedua nama foldernya
	router.Static("/images", "./images")
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/user/fetch", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campaignHandler.UploadCampaignImage)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), transctionHandler.GetCampaignTransctions)
	api.GET("/transactions", middleware.AuthMiddleware(authService, userService), transctionHandler.GetUserTransactions)
	api.POST("/transactions", middleware.AuthMiddleware(authService, userService), transctionHandler.CreateTransaction)
	api.POST("/transactions/noitification", transctionHandler.GetNotification)

	return router
}
