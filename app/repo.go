package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type AppRepo struct {
	db *pgxpool.Pool
}

func NewAppRepo(db *pgxpool.Pool) *AppRepo {
	return &AppRepo{
		db: db,
	}
}
