package dbrepo

import (
	"database/sql"

	"github.com/tinfante/web-app-learn/internal/config"
	"github.com/tinfante/web-app-learn/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
