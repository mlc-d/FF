package repository

import (
	"database/sql"
	"time"

	"gitlab.com/mlc-d/ff/pkg/entity"
)

type UserRepo interface {
	Register(u *entity.User) (*int64, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Register(u *entity.User) (*int64, error) {
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
