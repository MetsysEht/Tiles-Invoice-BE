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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "https://localhost:5173"}
	config.AllowCredentials = true
	S.Use(cors.New(config))
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
	userRouter.POST("/create", userServer.Create)
	userRouter.GET("/all", userServer.GetAll)
}
