package server

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/boot"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/users"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/logger"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

var S *gin.Engine

func Initialize() {
	gin.SetMode(gin.ReleaseMode)
	S = gin.New()
	S.Use(cors.Default())
	//S.Use(middleware.CheckAuthMiddleware)
	S.Use(ginzap.RecoveryWithZap(logger.L.Desugar(), true))

	registerRoutes()
}

func registerRoutes() {
	userRepo := users.NewRepo(boot.DB)
	userManager := users.NewManager(userRepo)
	userServer := users.CreateServer(userManager)

	S.POST("users/login", userServer.Login)
	userRouter := S.Group("/users")
	//userRouter.Use(middleware.AuthzMiddleware())
	userRouter.POST("/", userServer.Create)
	userRouter.GET("/", userServer.GetAll)
	userRouter.PUT("/", userServer.Update)
	userRouter.DELETE("/", userServer.Delete)
}
