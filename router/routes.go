package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/handler"
	"github.com/project-sistem-voucher/config"
	"github.com/project-sistem-voucher/manager"
	"github.com/project-sistem-voucher/middleware"
	"github.com/sirupsen/logrus"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middleware.LogRequestMiddleware(logrus.New()))

	rdb := config.NewCacher(*config.Cfg, 50)
	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager, rdb)

	middleware := middleware.NewMiddleware(rdb)

	voucherHandler := handler.NewVoucherHandler(repoManager.VoucherService())
	redeemHandler := handler.NewRedeemHandler(repoManager.RedeemService())
	userHandler := handler.NewUserHandler(repoManager.UserService())
	authHandler := handler.NewAuthHandler(repoManager.AuthService())

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/register", userHandler.RegisterUser)
		auth.POST("/login", authHandler.Login)
	}

	{
		sistemVoucher := v1.Group("/management-voucher")
		{
			sistemVoucher.Use(middleware.Authentication())
			sistemVoucher.POST("/create", voucherHandler.CreateVoucher)
			sistemVoucher.DELETE("/delete/:id", voucherHandler.DeleteVoucher)
			sistemVoucher.PUT("update/:id", voucherHandler.UpdateVoucher)
			sistemVoucher.GET("list", voucherHandler.GetVouchers)
			sistemVoucher.GET("/redeem-list", voucherHandler.GetVouchersForRedeem)
			sistemVoucher.POST("/redeem", redeemHandler.RedeemVoucher)
		}
	}

	return router.Run()

}
