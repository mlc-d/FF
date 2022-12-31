package internal

import (
	"database/sql"
	"gitlab.com/mlc-d/ff/db"
)

var (
	sqlDB = db.GetDB()
)

type TopicRepo interface {
	Post(t *Topic) (*int64, error)
}

type topicRepo struct {
	db *sql.DB
}

func NewTopicRepo() TopicRepo {
	return &topicRepo{
		db: sqlDB,
	}
}

const (
	post = `insert into topics (short_name, name, thumbnail_url, is_nsfw, maximum_threads, created_by, created_at) values (?, ?, ?, ?, ?, ?, ?);`
)

func (tr *topicRepo) Post(t *Topic) (*int64, error) {
	result, err := tr.db.Exec(post,
		t.ShortName,
		t.Name,
		t.ThumbnailURL,
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
