package repositories

import (
	"masivo/model"
)

type Placebo struct{}

func (repo *Placebo) Nombre() string {
	return "Placebo"
}

func (repo *Placebo) Inicializar() error {
	return nil
}

func (repo *Placebo) InsertarLote(registros []model.Registro) error {
	return nil
}

func (repo *Placebo) Limpiar() error {
	return nil
}

func (repo *Placebo) Cerrar() error {
	return nil
}
