package app

type AppData struct {
	DataId  int    `db:"data_id"`
	Message string `db:"message"`
}

type UserAppData interface {
	FindAppDataByID(ID int) (*AppData, error)
}

// FindByID ..
func (r *AppRepo) FindAppDataByID(ID int) (*AppData, error) {
	return &AppData{Message: "Hello World!"}, nil
}
