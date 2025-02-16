package interfaces

import "masivo/model"

type Repositorio interface {
	Nombre() string
	Inicializar() error
	InsertarLote(registros []model.Registro) error
	Limpiar() error
	Cerrar() error
}
