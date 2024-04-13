package routes

import (
	"avito/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupBannerRoutes(r *gin.Engine, bannerCtrl *controllers.BannerController) {
	r.GET("/user_banner", bannerCtrl.UserBanner)
	r.GET("/banner", bannerCtrl.GetBanner)
	r.POST("/banner", bannerCtrl.SaveBanner)
	r.PATCH("/banner/:id", bannerCtrl.UpdateBanner)
	r.DELETE("banner/:id", bannerCtrl.DeleteBanner)
}
