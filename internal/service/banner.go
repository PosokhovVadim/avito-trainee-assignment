package service

import (
	"avito/internal/model/controllermodel"
	"avito/internal/model/servicemodel"
	"avito/internal/storage"
	"avito/internal/utils"
	"avito/pkg/logger"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
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
	intParams, err := utils.ConvertToInt(tagID, featureID)

	if err != nil {
		b.log.Error("int convert failed", logger.Err(err))
		return nil, err
	}

	fl, err := strconv.ParseBool(useLastRevision)
	if err != nil {
		b.log.Error("bool convert failed", logger.Err(err))
		return nil, err
	}

	var bytes []byte
	if fl {
		bytes, err = b.s.GetUserBannerLastRevision(intParams[0], intParams[1])
	} else {
		bytes, err = b.s.GetUserBanner(intParams[0], intParams[1])
	}

	if err != nil {
		b.log.Error("select error", logger.Err(err))
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		b.log.Error("unmarshal error", logger.Err(err))
		return nil, err
	}

	return data, nil
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
