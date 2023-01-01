package internal

import (
	"database/sql"
	"gitlab.com/mlc-d/ff/db"
)

var (
	sqlDB = db.GetDB()
)

type Repo interface {
	Create(t *Topic) (*int64, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo() Repo {
	return &repo{
		db: sqlDB,
	}
}

const (
	post = `insert into topics (short_name, name, media_id, is_nsfw, maximum_threads, created_by, created_at) values (?, ?, ?, ?, ?, ?, ?);`
)

func (tr *repo) Create(t *Topic) (*int64, error) {
	result, err := tr.db.Exec(post,
		t.ShortName,
		t.Name,
		t.MediaID,
		t.IsNSFW,
		t.MaximumThreads,
		t.CreatedBy,
		t.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}
