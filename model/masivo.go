package model

import (
	"time"
)

// Registro representa un registro con los campos de la tabla registros
type Registro struct {
	Nombre      string    `json:"nombre"`
	Domicilio   string    `json:"domicilio"`
	Comentarios string    `json:"comentarios"`
	Puntaje     int       `json:"puntaje"`
	Ingreso     float64   `json:"ingreso"`
	Fecha       time.Time `json:"fecha"`
}
