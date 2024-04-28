package config

import (
	"rr-backend/config/db"
	"rr-backend/config/ent"
	"rr-backend/ent/entgen"

	"github.com/jmoiron/sqlx"
)

type IAppConfig interface {
	GetSqlxClient() *sqlx.DB
	GetENTClient() *entgen.Client
}

func NewAppConfig() IAppConfig {
	return &sAppConfig{
		SqlxClient: db.GetSqlxClient(),
		ENTClient:  ent.GetENTClient(),
	}
}

type sAppConfig struct {
	SqlxClient *sqlx.DB
	ENTClient  *entgen.Client
}

func (s *sAppConfig) GetSqlxClient() *sqlx.DB {
	return s.SqlxClient
}

func (s *sAppConfig) GetENTClient() *entgen.Client {
	return s.ENTClient
}
