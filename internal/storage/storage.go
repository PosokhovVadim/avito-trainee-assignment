package storage

type Storage interface {
	GetUserBanner(tagID int64, featureID int64) ([]byte, error)
	GetUserBannerLastRevision(tagID int64, featureID int64) ([]byte, error)
}
