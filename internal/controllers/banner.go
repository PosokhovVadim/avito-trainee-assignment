package controllers

import "github.com/gin-gonic/gin"

type Banner interface {
	EmptyFunc()
}

type BannerController struct {
	banner Banner
}

func NewBannerController(banner Banner) *BannerController {
	return &BannerController{
		banner: banner,
	}
}

func (b *BannerController) EmptyFunc() gin.HandlerFunc {
	
	return gin.HandlerFunc(func(c *gin.Context) {
		b.banner.EmptyFunc()
		c.JSON(200, gin.H{
			"message": "Banner is doing nothing",
		})
	})
}
