package handler

import (
	"database/sql"

	"gitlab.com/mlc-d/ff/db"
	user_repo "gitlab.com/mlc-d/ff/pkg/user/repository"
)

var (
	userRepo         = user_repo.NewUserRepo(sqlDB)
	sqlDB    *sql.DB = db.GetDB()
)
