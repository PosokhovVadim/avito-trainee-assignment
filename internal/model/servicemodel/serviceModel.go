package servicemodel

type Banner struct {
	ID        int64 `json:"id"`
	Content   byte  `json:"content"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	IsActive  bool  `json:"is_active"`
}

type BannerTagFeature struct {
	ID        int64 `json:"id"`
	BannerID  int64 `json:"banner_id"`
	TagID     int64 `json:"tag_id"`
	FeatureID int64 `json:"feature_id"`
}
