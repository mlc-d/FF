package internal

import (
	"database/sql"
	"gitlab.com/mlc-d/ff/db"
)

const (
	insert = `insert into comments (thread_id, user_id, tag, body, unique_id, is_op, color) values (?, ?, ?, ?, ?, ?, ?)`
)

var (
	sqlDB = db.GetDB()
)

type CommentRepo interface {
	Post(c *Comment) (*int64, error)
}

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepo() CommentRepo {
	return &commentRepo{
		db: sqlDB,
	}
}

func (cr *commentRepo) Post(c *Comment) (*int64, error) {
	result, err := cr.db.Exec(insert, c.ThreadID, c.UserID, c.Tag, c.Content, c.UniqueID, c.IsOP, c.Color)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}
