package media_repo

import (
	"database/sql"
	"log"
	"time"
)

const (
	blacklistQuery = `select is_blacklisted from media where hash = ?`
	insert         = `insert into media (hash, extension, created_at) values (?, ?, ?)`
)

type MediaRepo interface {
	IsBlacklisted(hash string) (bool, error)
	Insert(hash, ext string, createdAt time.Time) (*int64, error)
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
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	}
	return flag, nil
}

func (mr *mediaRepo) Insert(hash, extension string, createdAt time.Time) (*int64, error) {
	result, err := mr.db.Exec(insert, hash, extension, createdAt)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}
