package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewStorage(storagePath string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(storagePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (s *Postgres) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (s *Postgres) GetUserBanner(tagID int64, featureID int64) ([]byte, error) {
	return []byte(`{
		"func" : "GetUserBanner"
	}`), nil
}

func (s *Postgres) GetUserBannerLastRevision(tagID int64, featureID int64) ([]byte, error) {
	return []byte(`{
		"func" : "GetUserBannerLastRevision"
	}`), nil
}
