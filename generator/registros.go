package generator

import "masivo/model"

func GenerarRegistros(cantidad int) []model.Registro {
	registros := make([]model.Registro, cantidad)
	for i := 0; i < cantidad; i++ {
		registros[i] = generaRegistro()
	}
	return registros
}
