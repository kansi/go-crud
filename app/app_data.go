package app

import "github.com/jackc/pgx/v4/pgxpool"

type AppData struct {
	DataId  int    `db:"data_id"`
	Message string `db:"message"`
}

type AppDataModel struct {
	db *pgxpool.Pool
}

func NewAppDataModel(db *pgxpool.Pool) *AppDataModel {
	return &AppDataModel{
		db: db,
	}
}

type UserAppData interface {
	FindAppDataByID(ID int) (*AppData, error)
}

// FindByID ..
func (m *AppDataModel) FindAppDataByID(ID int) (*AppData, error) {
	return &AppData{Message: "Hello World!"}, nil
}
