package service

import (
	"avito/internal/model/controllermodel"
	"avito/internal/model/servicemodel"
	"avito/internal/storage"
	"fmt"
	"log/slog"
)

type BannerService struct {
	log *slog.Logger
	s   storage.Storage
}

func NewBannerService(log *slog.Logger, storage storage.Storage) *BannerService {
	return &BannerService{
		log: log,
		s:   storage,
	}
}

func (b *BannerService) UserBanner(tagID string, featureID string, useLastRevision string) (interface{}, error) {
	// tagidNum, err := strconv.Atoi(tagID)
	// if err != nil {
	// 	b.log.Error("convert failed", slog.String("err", err.Error()))
	// 	return nil, err
	// }

	// featureidNum, err := strconv.Atoi(featureID)
	// if err != nil {
	// 	b.log.Error("convert failed", slog.String("err", err.Error()))
	// 	return nil, err
	// }

	// fl, err := strconv.ParseBool(useLastRevision)
	// if err != nil {
	// 	b.log.Error("convert failed", logger.Err(err))
	// 	return nil, err
	// }

	// var bytes []byte
	// if fl {
	// 	bytes, err = b.s.GetUserBannerLastRevision(int64(tagidNum), int64(featureidNum))
	// } else {
	// 	bytes, err = b.s.GetUserBanner(int64(tagidNum), int64(featureidNum))
	// }

	// if err != nil {
	// 	b.log.Error("select error", logger.Err(err))
	// 	return nil, err
	// }

	// var data interface{}
	// err = json.Unmarshal(bytes, &data)
	// if err != nil {
	// 	b.log.Error("unmarshal error", logger.Err(err))
	// 	return nil, err
	// }

	return nil, nil
}

func (b *BannerService) GetBanners(tagID string, featureID string, limit string, offset string) ([]servicemodel.Banner, error) {
	return nil, nil
}

func (b *BannerService) SaveBanner(ctrlBanner *controllermodel.Banner) (int64, error) {
	return -1, nil
}

func (b *BannerService) UpdateBanner(bannerID string, ctrlBanner *controllermodel.Banner) error {
	fmt.Println(*ctrlBanner.Content)

	return nil
}

func (b *BannerService) DeleteBanner(id string) error {
	return nil
}
