package service

import (
	"avito/internal/model/controllermodel"
	"avito/internal/model/servicemodel"
	"avito/internal/storage"
	"avito/pkg/logger"
	"encoding/json"
	"log/slog"
	"strconv"
	"time"
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

func (b *BannerService) parseNotRequiredInt(value string) (int64, error) {
	if value == "" {
		return storage.InvalidInt, nil
	}
	num, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		b.log.Error("convert failed", logger.Err(err))

		return 0, err
	}
	return num, nil
}

func (b *BannerService) UserBanner(tagID string, featureID string, useLastRevision string, adminStatus bool) (map[string]interface{}, error) {
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

	var banner *servicemodel.Banner
	if fl {
		banner, err = b.s.GetLastRevision(int64(tagidNum), int64(featureidNum), fl)
	} else {
		banner, err = b.s.GetBannerByTagAndFeature(int64(tagidNum), int64(featureidNum))
	}

	if err != nil {
		b.log.Error("postgres err: get banner failed", logger.Err(err))
		return nil, err
	}

	if !banner.IsActive && !adminStatus {
		return nil, storage.BannerNotFound
	}

	var jsonContent map[string]interface{}
	json.Unmarshal(banner.Content, &jsonContent)
	return jsonContent, nil
}

func (b *BannerService) GetBanners(tagID string, featureID string, limit string, offset string) (*[]servicemodel.Banner, error) {
	tagNum, err := b.parseNotRequiredInt(tagID)
	if err != nil {
		return nil, err
	}

	featureNum, err := b.parseNotRequiredInt(featureID)
	if err != nil {
		return nil, err
	}

	limitNum, err := b.parseNotRequiredInt(limit)
	if err != nil {
		return nil, err
	}

	offsetNum, err := b.parseNotRequiredInt(offset)
	if err != nil {
		return nil, err
	}
	banners, err := b.s.GetBanners(tagNum, featureNum, limitNum, offsetNum)
	if err != nil {
		return nil, err
	}
	return banners, nil
}

func (b *BannerService) SaveBanner(ctrlBanner *controllermodel.Banner) (int64, error) {
	byteContent, err := json.Marshal(ctrlBanner.Content)
	if err != nil {
		return storage.InvalidInt, err
	}

	bannerID, err := b.s.InsertBanner(&servicemodel.Banner{
		Content:   byteContent,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  *ctrlBanner.IsActive,
	})

	if err != nil {
		return storage.InvalidInt, err
	}

	if ctrlBanner.FeatureID != nil {
		_, err := b.s.InsertBannerFeature(&servicemodel.BannerFeature{BannerID: bannerID, FeatureID: *ctrlBanner.FeatureID})
		if err != nil {
			return storage.InvalidInt, err
		}
	}

	if ctrlBanner.TagID != nil {
		for _, v := range *ctrlBanner.TagID {
			_, err := b.s.InsertBannerTag(&servicemodel.BannerTag{BannerID: bannerID, TagID: v})
			if err != nil {
				return storage.InvalidInt, err
			}
		}
	}

	return bannerID, nil
}

func (b *BannerService) UpdateBanner(bannerID string, ctrlBanner *controllermodel.Banner) error {
	id, err := strconv.Atoi(bannerID)
	if err != nil {
		b.log.Error("convert failed", slog.String("err", err.Error()))
		return err
	}

	if ctrlBanner.Content != nil {
		byteContent, err := json.Marshal(ctrlBanner.Content)
		if err != nil {
			return err
		}
		err = b.s.UpdateBannerContent(int64(id), byteContent)
		if err != nil {
			return err
		}
	}

	if ctrlBanner.IsActive != nil {
		err = b.s.UpdateBannerActivity(int64(id), *ctrlBanner.IsActive)
		if err != nil {
			return err
		}
	}

	if ctrlBanner.FeatureID != nil {
		err = b.s.UpdateBannerFeature(int64(id), *ctrlBanner.FeatureID)
		if err != nil {
			return err
		}
	}

	if ctrlBanner.TagID != nil {
		bt, err := b.s.GetBannerTags(int64(id))
		if err != nil {
			return err
		}
		for _, v := range *bt {
			err = b.s.DeleteBannerTag(int64(id), v.TagID)
			if err != nil {
				return err
			}
		}
		for _, v := range *ctrlBanner.TagID {
			_, err := b.s.InsertBannerTag(&servicemodel.BannerTag{BannerID: int64(id), TagID: v})
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func (b *BannerService) DeleteBanner(id string) error {
	bannerID, err := strconv.Atoi(id)
	if err != nil {
		b.log.Error("convert failed", slog.String("err", err.Error()))
		return err
	}
	return b.s.DeleteBanner(int64(bannerID))
}
