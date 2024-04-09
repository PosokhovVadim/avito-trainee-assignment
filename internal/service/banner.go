package service

import (
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

func (b *BannerService) EmptyFunc() {
	fmt.Println("I am(service) doing nothing too")
	b.s.StorageEmptyFunc()
}
