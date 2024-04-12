package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Banner interface {
	UserBanner(tagID string, featureID string, useLastRevision string) (interface{}, error)
}

type BannerController struct {
	banner Banner
}

func NewBannerController(banner Banner) *BannerController {
	return &BannerController{
		banner: banner,
	}
}

func (b *BannerController) UserBanner(c *gin.Context) {
	tagID := c.Query("tag_id")
	featureID := c.Query("feature_id")
	if tagID == "" || featureID == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	useLastRevision := c.DefaultQuery("use_last_revision", "false")

	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, nil)
	}

	data, err := b.banner.UserBanner(tagID, featureID, useLastRevision)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, data)
}

func (b *BannerController) GetBanner(c *gin.Context) {
	//
}

func (b *BannerController) SaveBanner(c *gin.Context) {
	//
}

func (b *BannerController) UpdateBanner(c *gin.Context) {
	//
}

func (b *BannerController) DeleteBanner(c *gin.Context) {
	//
}
