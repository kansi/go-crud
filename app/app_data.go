package app

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

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

func (m *AppDataModel) StoreMessage(message string) (int, error) {
	return exec_query(m, "INSERT INTO app_data (message) VALUES ($1) RETURNING data_id", message)
}

func exec_query(m *AppDataModel, query string, arguments ...interface{}) (int, error) {
	id := -1

	ctx := context.Background()
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return id, err
	}

	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	// defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, query, arguments...).Scan(&id)
	if err != nil {
		return id, err
	}

	err = tx.Commit(ctx)

	return id, err
}
