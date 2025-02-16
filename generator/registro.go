package generator

import "masivo/model"

func generaRegistro() model.Registro {
	return model.Registro{
		Nombre:      GeneraCadena(50),
		Domicilio:   GeneraCadena(300),
		Comentarios: GeneraCadena(500),
		Puntaje:     generarEnteroAleatorio(),
		Ingreso:     generarNumeroReal(),
		Fecha:       generarFechaAleatoria(),
	}
}
