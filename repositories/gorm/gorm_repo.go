package gorm

import (
	"masivo/interfaces"
	"masivo/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormRepo struct {
	interfaces.Repositorio
	db *gorm.DB
}

func (repo *GormRepo) Nombre() string {
	return "GORM"
}

func (repo *GormRepo) Inicializar() error {
	dsn := "user=masivo password=Ap2485!aO65>AV_o8 host=127.0.0.1 port=5432 dbname=masivo sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}

func (repo *GormRepo) InsertarLote(registros []model.Registro) error {
	return repo.db.Create(&registros).Error
}

func (repo *GormRepo) Limpiar() error {
	return repo.db.Exec("TRUNCATE TABLE registros").Error
}

func (repo *GormRepo) Cerrar() error {
	sqlDB, err := repo.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
