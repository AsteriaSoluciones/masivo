package main

import (
	"fmt"
	"masivo/generator"
	"masivo/interfaces"
	"masivo/repositories"
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
func mostrarAyuda() {
	fmt.Println("Uso de la aplicación:")
	fmt.Println("  masivo <tipo_repositorio>")
	fmt.Println("Tipos de repositorio soportados:")
	fmt.Println("  sqlite, placebo, pgx, gorm, mongo, mysql")
}

func main() {
	fmt.Println("Inserción Masiva")

	if len(os.Args) < 2 {
		mostrarAyuda()
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
		repo = &repositories.SqliteRepo{}
	case "placebo":
		repo = &repositories.Placebo{}
	case "pgx":
		repo = &repositories.PgxRepo{}
	case "gorm":
		repo = &repositories.GormRepo{}
	case "mongo":
		repo = &repositories.Mongo{}
	case "mysql":
		repo = &repositories.MysqlRepo{}
	default:
		fmt.Println("Tipo de repositorio no soportado")
		return
	}

	proceso(repo, tamanoLote, cantidadLotes)
}
