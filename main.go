package main

import (
	"fmt"
	"masivo/generator"
	"masivo/interfaces"
	"masivo/repositories/mongo"
	"masivo/repositories/mysql"
	"masivo/repositories/pgx"
	"masivo/repositories/placebo"
	"masivo/repositories/sqlite"
	"os"
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

	if len(os.Args) < 2 {
		fmt.Println("Debe especificar el tipo de repositorio")
		return
	}

	tipoRepo := os.Args[1]
	tamanoLote := 1000
	cantidadLotes := 1000
	total := tamanoLote * cantidadLotes
	fmt.Println("Total de registros a insertar: ", total)

	var repo interfaces.Repositorio

	switch tipoRepo {
	case "sqlite":
		repo = &sqlite.SqliteRepo{}
	case "placebo":
		repo = &placebo.Placebo{}
	case "pgx":
		repo = &pgx.PgxRepo{}
	case "mongo":
		repo = &mongo.Mongo{}
	case "mysql":
		repo = &mysql.MysqlRepo{}
	default:
		fmt.Println("Tipo de repositorio no soportado")
		return
	}

	proceso(repo, tamanoLote, cantidadLotes)
}
