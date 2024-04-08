package app

import (
	"avito/internal/controllers"
	"avito/internal/routes"
	"avito/internal/service"
	"avito/pkg/logger"

	"log/slog"

	"github.com/gin-gonic/gin"
)

type App struct {
	log     *slog.Logger
	address string
	Gin     *gin.Engine
}

func NewBannerApp(log *slog.Logger, address string, storagePath string) *App {
	_ = storagePath

	r := ginSetup()
	bannerService := service.NewBannerService(log)

	bannerCtrl := controllers.NewBannerController(bannerService)

	routes.SetupBannerRoutes(r, bannerCtrl)
	return &App{
		log:     log,
		address: address,
		Gin:     r,
	}
}

func ginSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	return r
}

func (a *App) Run() error {
	a.log.Info("Starting http server:", slog.String("addr", a.address))
	err := a.Gin.Run(a.address)
	if err != nil {
		a.log.Error("Cannot start server:", logger.Err(err))
		return err
	}
	return nil
}
