package config

import (
	"rr-backend/config/db"

	"github.com/jmoiron/sqlx"
)

type IAppConfig interface {
	GetSqlxClient() *sqlx.DB
}

func NewAppConfig() IAppConfig {
	return &sAppConfig{
		SqlxClient: db.GetSqlxClient(),
	}
}

type sAppConfig struct {
	SqlxClient *sqlx.DB
}

func (s *sAppConfig) GetSqlxClient() *sqlx.DB {
	return s.SqlxClient
}
