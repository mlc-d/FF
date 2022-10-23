package media_repo

import "database/sql"

const (
	blacklistQuery = `select is_blacklisted from media where hash = ?`
)

type MediaRepo interface {
	IsBlacklisted(hash string) (bool, error)
}

type mediaRepo struct {
	db *sql.DB
}

func NewMediaRepo(db *sql.DB) MediaRepo {
	return &mediaRepo{
		db: db,
	}
}

func (mr *mediaRepo) IsBlacklisted(hash string) (bool, error) {
	var flag bool
	err := mr.db.QueryRow(blacklistQuery, hash).
		Scan(&flag)
	return flag, err
}
