package thread_repo

import (
	"database/sql"

	"gitlab.com/mlc-d/ff/db"
	"gitlab.com/mlc-d/ff/pkg/thread"
)

var (
	sqlDB = db.GetDB()
)

type ThreadRepo interface {
	Post(t *thread.Thread) (*int64, error)
}

type threadRepo struct {
	db *sql.DB
}

func NewThreadRepo() ThreadRepo {
	return &threadRepo{
		db: sqlDB,
	}
}

const (
	post = `insert into threads (topic_id, user_id, hash, title, body, media_id, sticky, created_at) values (?, ?, ?, ?, ?, ?, ?, ?);`
)

func (tr *threadRepo) Post(t *thread.Thread) (*int64, error) {
	result, err := tr.db.Exec(post,
		t.TopicID,
		t.UserID,
		t.Hash,
		t.Title,
		t.Body,
		t.MediaID,
		t.Sticky,
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
