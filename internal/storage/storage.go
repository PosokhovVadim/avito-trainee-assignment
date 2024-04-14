package storage

import (
	"avito/internal/model/servicemodel"
	"fmt"
)

type Storage interface {
	GetBannerByTagAndFeature(tagId, featureId int64) (*servicemodel.Banner, error)
	GetLastRevision(tagId, featureId int64, lastRevision bool) (*servicemodel.Banner, error)
	GetBanners(tagId, featureId, limit, offset int64) (*[]servicemodel.Banner, error)

	GetBannerTags(bannerID int64) (*[]servicemodel.BannerTag, error)
	GetBannerFeature(bannerID int64) (*servicemodel.BannerFeature, error)

	InsertBanner(banner *servicemodel.Banner) (int64, error)
	InsertBannerFeature(bannerFeature *servicemodel.BannerFeature) (int64, error)
	InsertBannerTag(bannerTag *servicemodel.BannerTag) (int64, error)

	UpdateBannerContent(id int64, content []byte) error
	UpdateBannerActivity(id int64, isActive bool) error
	UpdateBannerFeature(id int64, featureID int64) error

	DeleteBanner(id int64) error
	DeleteBannerTag(bannerID, tagID int64) error
	DeleteBannerFeature(bannerID, featureID int64) error
}

const (
	InvalidInt = -1
)

func ConditionalSQL(value int64, format string) string {
	if value != InvalidInt {
		return fmt.Sprintf(" "+format, value)
	}
	return ""
}
