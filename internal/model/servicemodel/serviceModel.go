package servicemodel

import "time"

type Banner struct {
	ID        int64     `json:"id"`
	Content   []byte    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

type BannerTag struct {
	ID       int64 `json:"id"`
	BannerID int64 `json:"banner_id"`
	TagID    int64 `json:"tag_id"`
}

type BannerFeature struct {
	ID        int64 `json:"id"`
	BannerID  int64 `json:"banner_id"`
	FeatureID int64 `json:"feature_id"`
}
