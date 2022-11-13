package user_repo

import (
	"database/sql"
	"time"

	"gitlab.com/mlc-d/ff/db"
	"gitlab.com/mlc-d/ff/pkg/user"
)

var (
	sqlDB = db.GetDB()
)

type UserRepo interface {
	Register(u *user.User) (*int64, *uint8, error)
	GetPassword(nick string) (*int64, *uint8, string, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo() UserRepo {
	return &userRepo{
		db: sqlDB,
	}
}

func (ur *userRepo) Register(u *user.User) (*int64, *uint8, error) {
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

func (ur *userRepo) GetPassword(nick string) (*int64, *uint8, string, error) {
	var passwordFromDB string
	var id int64
	var role uint8
	err := ur.db.QueryRow(`select id, role_id, password from users where nick = ?`, nick).
		Scan(&id, &role, &passwordFromDB)
	return &id, &role, passwordFromDB, err
}
