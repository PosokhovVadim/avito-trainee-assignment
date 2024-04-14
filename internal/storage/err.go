package storage

import "errors"

var (
	BannerNotFound      = errors.New("Banner not found")
	BannerAlreadyExists = errors.New("Banner already exists")
)
