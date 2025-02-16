package placebo

import (
	"masivo/interfaces"
	"masivo/model"
)

type Placebo struct {
	interfaces.Repositorio
}

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
