package pgx

import (
	"context"
	"masivo/interfaces"
	"masivo/model"

	"github.com/jackc/pgx/v4"
)

type PgxRepo struct {
	interfaces.Repositorio
	conn *pgx.Conn
}

func (repo *PgxRepo) Nombre() string {
	return "PGX"
}

func (repo *PgxRepo) Inicializar() error {
	connNueva, err := pgx.Connect(context.Background(), "user=masivo password=Ap2485!aO65>AV_o8 host=127.0.0.1 port=5432 dbname=masivo sslmode=disable")

	if err == nil {
		repo.conn = connNueva
	}

	return err
}

func (repo *PgxRepo) InsertarLote(registros []model.Registro) error {

	// Usar CopyFrom para realizar el bulk insert
	// `CopyFrom` es mucho más rápido que realizar múltiples `INSERT` individuales.
	_, err := repo.conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"registros"}, // Nombre de la tabla
		[]string{"nombre", "domicilio", "comentarios", "puntaje", "ingreso", "fecha"}, // Columnas a insertar
		pgx.CopyFromRows(convertRegistrosToRows(registros)),                           // Función que convierte a filas para COPY
	)

	return err
}

func (repo *PgxRepo) Limpiar() error {
	_, err := repo.conn.Exec(context.Background(), "TRUNCATE TABLE registros")
	return err
}

func (repo *PgxRepo) Cerrar() error {
	return repo.conn.Close(context.Background())
}

// Convierte un slice de registros a un slice de `[]interface{}`
// que es lo que `CopyFrom` necesita
func convertRegistrosToRows(registros []model.Registro) [][]interface{} {
	var rows [][]interface{}
	for _, registro := range registros {
		rows = append(rows, []interface{}{registro.Nombre, registro.Domicilio, registro.Comentarios, registro.Puntaje, registro.Ingreso, registro.Fecha})
	}
	return rows
}
