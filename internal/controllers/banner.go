package controllers

import (
	"avito/internal/controllers/permission"
	"avito/internal/model/controllermodel"
	"avito/internal/model/servicemodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Banner interface {
	UserBanner(tagID string, featureID string, useLastRevision string) (interface{}, error)
	GetBanners(tagID string, featureID string, limit string, offset string) ([]servicemodel.Banner, error)
	SaveBanner(ctrlBanner *controllermodel.Banner) (int64, error)
	UpdateBanner(bannerID string, ctrlBanner *controllermodel.Banner) error
	DeleteBanner(bannerID string) error
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

	token := c.GetHeader("token")
	if !permission.IsVadlidToken(token) {
		c.JSON(http.StatusUnauthorized, Unauthorized)
		return
	}

	tagID := c.Query("tag_id")
	featureID := c.Query("feature_id")
	if tagID == "" || featureID == "" {
		c.JSON(http.StatusBadRequest, BadRequest)
		return
	}

	useLastRevision := c.DefaultQuery("use_last_revision", "false")

	data, err := b.banner.UserBanner(tagID, featureID, useLastRevision)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalError)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (b *BannerController) GetBanner(c *gin.Context) {

	token := c.GetHeader("token")
	if !permission.IsVadlidToken(token) {
		c.JSON(http.StatusUnauthorized, Unauthorized)
		return
	}

	if !permission.IsAdmin(token) {
		c.JSON(http.StatusForbidden, AccessDenied)
		return
	}

	tagID := c.Query("tag_id")
	featureID := c.Query("feature_id")
	limit := c.Query("limit")
	offset := c.Query("offset")

	banners, err := b.banner.GetBanners(tagID, featureID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalError)
		return
	}

	c.JSON(http.StatusOK, banners)
}

func (b *BannerController) SaveBanner(c *gin.Context) {
	token := c.GetHeader("token")
	if !permission.IsVadlidToken(token) {
		c.JSON(http.StatusUnauthorized, Unauthorized)
		return
	}

	if !permission.IsAdmin(token) {
		c.JSON(http.StatusForbidden, AccessDenied)
		return
	}

	ctrlBanner := &controllermodel.Banner{}
	if err := c.BindJSON(ctrlBanner); err != nil {
		c.JSON(http.StatusBadRequest, BadRequest)
		return
	}

	if id, err := b.banner.SaveBanner(ctrlBanner); err != nil {
		c.JSON(http.StatusInternalServerError, InternalError)
		return
	} else {
		c.JSON(http.StatusCreated, id)
	}
}

func (b *BannerController) UpdateBanner(c *gin.Context) {

	token := c.GetHeader("token")
	if !permission.IsVadlidToken(token) {
		c.JSON(http.StatusUnauthorized, Unauthorized)
		return
	}

	if !permission.IsAdmin(token) {
		c.JSON(http.StatusForbidden, AccessDenied)
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, BadRequest)
		return
	}

	ctrlBanner := &controllermodel.Banner{}
	if err := c.BindJSON(ctrlBanner); err != nil {
		c.JSON(http.StatusBadRequest, BadRequest)
		return
	}

	if err := b.banner.UpdateBanner(id, ctrlBanner); err != nil {
		//todo: add custom db error handling
		c.JSON(http.StatusInternalServerError, InternalError)
		return
	}

	c.Status(http.StatusOK)

}

func (b *BannerController) DeleteBanner(c *gin.Context) {
	token := c.GetHeader("token")
	if !permission.IsVadlidToken(token) {
		c.JSON(http.StatusUnauthorized, Unauthorized)
		return
	}

	if !permission.IsAdmin(token) {
		c.JSON(http.StatusForbidden, AccessDenied)
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, BadRequest)
		return
	}

	if err := b.banner.DeleteBanner(id); err != nil {
		c.JSON(http.StatusInternalServerError, InternalError)
		return
	}

	c.Status(http.StatusNoContent)
}
