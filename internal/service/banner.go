package service

import (
	"avito/internal/storage"
	"avito/pkg/logger"
	"encoding/json"
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
	tagidNum, err := strconv.Atoi(tagID)
	if err != nil {
		b.log.Error("convert failed", slog.String("err", err.Error()))
		return nil, err
	}

	featureidNum, err := strconv.Atoi(featureID)
	if err != nil {
		b.log.Error("convert failed", slog.String("err", err.Error()))
		return nil, err
	}

	fl, err := strconv.ParseBool(useLastRevision)
	if err != nil {
		b.log.Error("convert failed", logger.Err(err))
		return nil, err
	}

	var bytes []byte
	if fl {
		bytes, err = b.s.GetUserBannerLastRevision(int64(tagidNum), int64(featureidNum))
	} else {
		bytes, err = b.s.GetUserBanner(int64(tagidNum), int64(featureidNum))
	}

	if err != nil {
		b.log.Error("select error", logger.Err(err))
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
