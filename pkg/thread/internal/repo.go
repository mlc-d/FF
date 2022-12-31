package internal

import (
	"database/sql"
	"gitlab.com/mlc-d/ff/db"
)

var (
	sqlDB = db.GetDB()
)

type Repo interface {
	Post(t *Thread) (*int64, error)
	GetUserID(int64) *int64
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
	post      = `insert into threads (topic_id, user_id, hash, title, body, media_id, sticky, created_at) values (?, ?, ?, ?, ?, ?, ?, ?);`
	getUserID = `select user_id from threads where id = ?;`
)

func (tr *repo) Post(t *Thread) (*int64, error) {
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

func (tr *repo) GetUserID(id int64) *int64 {
	var userID int64
	row := tr.db.QueryRow(getUserID, id)
	err := row.Scan(&userID)
	if err != nil {
		return nil
	}
	return &userID
}
