package main

import (
	"fmt"
	"masivo/generator"
	"masivo/interfaces"
	"masivo/repositories/sqlite"
)

func proceso(rep interfaces.Repositorio, tamanoLote int, cantidadLotes int) {
	fmt.Println("Iniciando proceso")
	fmt.Println("Repositorio: ", rep.Nombre())

	err := rep.Inicializar()
	if err != nil {
		fmt.Println("Error al conectar: ", err.Error())
		return
	}

	for i := 0; i < cantidadLotes; i++ {
		registros := generator.GenerarRegistros(tamanoLote)
		err = rep.InsertarLote(registros)
		if err != nil {
			fmt.Println("Error al insertar: ", err.Error())
			return
		}
	}

	err = rep.Limpiar()
	if err != nil {
		fmt.Println("Error al limpiar: ", err.Error())
		return
	}

	err = rep.Cerrar()
	if err != nil {
		fmt.Println("Error al cerrar: ", err.Error())
		return
	}

	fmt.Println("Proceso finalizado")
}

func main() {
	fmt.Println("Inserción Masiva")

	tamanoLote := 1000
	cantidadLotes := 1000
	total := tamanoLote * cantidadLotes
	fmt.Println("Total de registros a insertar: ", total)

	//proceso(&placebo.Placebo{}, tamanoLote, cantidadLotes)
	//proceso(&pgx.PgxRepo{}, tamanoLote, cantidadLotes)
	//proceso(&mongo.Mongo{}, tamanoLote, cantidadLotes)
	//proceso(&mysql.MysqlRepo{}, tamanoLote, cantidadLotes)
	proceso((&sqlite.SqliteRepo{}), tamanoLote, cantidadLotes)
}
