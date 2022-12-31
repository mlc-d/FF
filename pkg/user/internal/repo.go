package internal

import (
	"database/sql"
	"time"

	"gitlab.com/mlc-d/ff/db"
)

var (
	sqlDB = db.GetDB()
)

type Repo interface {
	Register(u *User) (*int64, *uint8, error)
	GetPassword(nick string) (*int64, *uint8, string, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo() Repo {
	return &repo{
		db: sqlDB,
	}
}

func (ur *repo) Register(u *User) (*int64, *uint8, error) {
	res, err := ur.db.Exec(`insert into users (nick, password, role_id, created_at) values (?, ?, ?, ?)`,
		u.Nick,
		u.Password,
		u.RoleID,
		time.Now().UTC(),
	)
	if err != nil {
		return nil, nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, nil, err
	}
	return &id, &u.RoleID, nil
}

func (ur *repo) GetPassword(nick string) (*int64, *uint8, string, error) {
	var passwordFromDB string
	var id int64
	var role uint8
	err := ur.db.QueryRow(`select id, role_id, password from users where nick = ?`, nick).
		Scan(&id, &role, &passwordFromDB)
	return &id, &role, passwordFromDB, err
}
