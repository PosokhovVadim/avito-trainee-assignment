package routes

import (
	"avito/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupBannerRoutes(r *gin.Engine, bannerCtrl *controllers.BannerController) {
	// banner := r.Group("/banner")
	// {
	// 	banner.GET("/", bannerCtrl.EmptyFunc())
	// }
	// _ = banner
	r.GET("/banner", bannerCtrl.EmptyFunc())
}
