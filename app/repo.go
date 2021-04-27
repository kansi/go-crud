package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type AppRepo struct {
	// Register models in the Repo so that they can be queried via DB pool
	AppData *AppDataModel
}

func NewAppRepo(dbConnPool *pgxpool.Pool) *AppRepo {
	appData := NewAppDataModel(dbConnPool)

	return &AppRepo{
		AppData: appData,
	}
}
