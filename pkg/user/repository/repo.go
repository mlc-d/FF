package user_repo

import (
	"database/sql"
	"time"

	"gitlab.com/mlc-d/ff/pkg/user"
)

type UserRepo interface {
	Register(u *user.User) (*int64, error)
	GetPassword(nick string) (string, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Register(u *user.User) (*int64, error) {
	res, err := ur.db.Exec(`insert into users (nick, password, role_id, created_at) values (?, ?, ?, ?)`,
		u.Nick,
		u.Password,
		u.RoleID,
		time.Now().UTC(),
	)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (ur *userRepo) GetPassword(nick string) (string, error) {
	var passwordFromDB string
	err := ur.db.QueryRow(`select password from users where nick = ?`, nick).
		Scan(&passwordFromDB)
	if err != nil {
		return "", err
	}
	return passwordFromDB, nil
}
