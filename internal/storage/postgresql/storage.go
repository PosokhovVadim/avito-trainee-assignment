package postgresql

import (
	"avito/internal/model/servicemodel"
	"avito/internal/storage"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Postgres, error) {
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (s *Postgres) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *Postgres) GetBannerByTagAndFeature(tagID, featureID int64) (*servicemodel.Banner, error) {
	tagSubquery := `SELECT banner_id FROM banner_tag WHERE tag_id = $1`
	featureSubquery := `SELECT banner_id FROM banner_feature WHERE feature_id = $2`
	query := `
        SELECT *
        FROM banner
        WHERE id IN (` + tagSubquery + `) AND id IN (` + featureSubquery + `)
		LIMIT 1
    `

	row := s.db.QueryRowContext(context.Background(), query, tagID, featureID)

	var banner servicemodel.Banner
	if err := row.Scan(
		&banner.ID,
		&banner.Content,
		&banner.CreatedAt,
		&banner.UpdatedAt,
		&banner.IsActive,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.BannerNotFound
		}

		return nil, err
	}
	return &banner, nil
}

func (s *Postgres) GetLastRevision(tagID, featureID int64, lastRevision bool) (*servicemodel.Banner, error) {
	tagSubquery := `SELECT banner_id FROM banner_tag WHERE tag_id = $1`
	featureSubquery := `SELECT banner_id FROM banner_feature WHERE feature_id = $2`
	query := `
        SELECT banner.*
        FROM banner
        WHERE id IN (` + tagSubquery + `) AND id IN (` + featureSubquery + `)
        ORDER BY updated_at DESC
        LIMIT 1
    `
	row := s.db.QueryRowContext(context.Background(), query, tagID, featureID)
	var banner servicemodel.Banner
	if err := row.Scan(
		&banner.ID,
		&banner.Content,
		&banner.CreatedAt,
		&banner.UpdatedAt,
		&banner.IsActive,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.BannerNotFound
		}

		return nil, err
	}
	return &banner, nil
}

func (s *Postgres) GetBanners(tagID, featureID, limit, offset int64) (*[]servicemodel.Banner, error) {

	tagSubquery := `SELECT banner_id FROM banner_tag WHERE 1=1 ` + storage.ConditionalSQL(tagID, " AND tag_id = %d")
	featureSubquery := `SELECT banner_id FROM banner_feature WHERE 1=1 ` + storage.ConditionalSQL(featureID, " AND feature_id = %d")

	sql := `SELECT * FROM banner WHERE id IN (` + tagSubquery + `) AND id IN (` + featureSubquery + `)` +
		storage.ConditionalSQL(limit, " LIMIT %d") +
		storage.ConditionalSQL(offset, " OFFSET %d")

	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []servicemodel.Banner
	for rows.Next() {
		var banner servicemodel.Banner
		if err := rows.Scan(
			&banner.ID,
			&banner.Content,
			&banner.CreatedAt,
			&banner.UpdatedAt,
			&banner.IsActive,
		); err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}
	return &banners, nil
}

func (s *Postgres) GetBannerTags(bannerID int64) (*[]servicemodel.BannerTag, error) {
	query := "SELECT * FROM banner_tag WHERE banner_id = $1"
	rows, err := s.db.QueryContext(context.Background(), query, bannerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bannerTags []servicemodel.BannerTag
	for rows.Next() {
		var bannerTag servicemodel.BannerTag
		if err := rows.Scan(
			&bannerTag.ID,
			&bannerTag.BannerID,
			&bannerTag.TagID,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, storage.BannerNotFound
			}
			return nil, err
		}
		bannerTags = append(bannerTags, bannerTag)
	}

	return &bannerTags, nil
}

func (s *Postgres) GetBannerFeature(bannerID int64) (*servicemodel.BannerFeature, error) {
	query := "SELECT * FROM banner_feature WHERE banner_id = $1"
	row := s.db.QueryRowContext(context.Background(), query, bannerID)
	var bannerFeature servicemodel.BannerFeature
	if err := row.Scan(
		&bannerFeature.ID,
		&bannerFeature.BannerID,
		&bannerFeature.FeatureID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.BannerNotFound
		}
		return nil, err
	}
	return &bannerFeature, nil
}

func (s *Postgres) InsertBanner(banner *servicemodel.Banner) (int64, error) {
	query := `
        INSERT INTO banner (content, created_at, updated_at, is_active)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var id int64
	if err := s.db.QueryRowContext(context.Background(), query, banner.Content, banner.CreatedAt, banner.UpdatedAt, banner.IsActive).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Postgres) InsertBannerFeature(banner *servicemodel.BannerFeature) (int64, error) {
	query := `
		INSERT INTO banner_feature (banner_id, feature_id)
		VALUES ($1, $2)
		RETURNING id
	`
	var id int64
	if err := s.db.QueryRowContext(context.Background(), query, banner.BannerID, banner.FeatureID).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Postgres) InsertBannerTag(banner *servicemodel.BannerTag) (int64, error) {
	query := `
		INSERT INTO banner_tag (banner_id, tag_id)
		VALUES ($1, $2)
		RETURNING id
	`
	var id int64
	if err := s.db.QueryRowContext(context.Background(), query, banner.BannerID, banner.TagID).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Postgres) UpdateBannerContent(id int64, content []byte) error {
	_, err := s.db.ExecContext(
		context.Background(),
		"UPDATE banner SET content = $1, updated_at = now() WHERE id = $2",
		content,
		id,
	)
	return err
}

func (s *Postgres) UpdateBannerActivity(id int64, isActive bool) error {
	_, err := s.db.ExecContext(
		context.Background(),
		"UPDATE banner SET is_active = $1, updated_at = now() WHERE id = $2",
		isActive,
		id,
	)
	return err
}

func (s *Postgres) UpdateBannerFeature(bannerID int64, featureID int64) error {
	_, err := s.db.ExecContext(context.Background(), "UPDATE banner_feature SET feature_id = $1 WHERE banner_id = $2", featureID, bannerID)
	return err
}

func (s *Postgres) DeleteBanner(id int64) error {
	res, err := s.db.ExecContext(context.Background(), "DELETE FROM banner WHERE id = $1", id)
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	} else if rowsAffected == 0 {
		return storage.BannerNotFound
	}

	return err
}

func (s *Postgres) DeleteBannerTag(bannerID, tagID int64) error {
	res, err := s.db.ExecContext(context.Background(), "DELETE FROM banner_tag WHERE banner_id = $1 AND tag_id = $2", bannerID, tagID)
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	} else if rowsAffected == 0 {
		return storage.BannerNotFound
	}

	return err
}

func (s *Postgres) DeleteBannerFeature(bannerID, featureID int64) error {
	res, err := s.db.ExecContext(context.Background(), "DELETE FROM banner_feature WHERE banner_id = $1 AND feature_id = $2", bannerID, featureID)
	if rowsAffected, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	} else if rowsAffected == 0 {
		return storage.BannerNotFound
	}
	return err
}
