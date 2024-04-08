package service

import (
	"fmt"
	"log/slog"
)

type BannerService struct {
	log *slog.Logger
}

func NewBannerService(log *slog.Logger) *BannerService {
	return &BannerService{
		log: log,
	}
}

func (b *BannerService) EmptyFunc() {
	fmt.Println("I am doing nothing too")
}
