package postgresql

import (
	"fmt"

	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB // do i need logger here?
}

func NewStorage(storagePath string) (*Postgres, error) {
	// db, err := gorm.Open(postgres.Open(storagePath), &gorm.Config{})
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
	// return &Postgres{db: db}, nil
}

func (s *Postgres) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (s *Postgres) StorageEmptyFunc() {
	fmt.Println("Aboba")
}
