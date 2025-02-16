package sqlite

import (
	"database/sql"
	"log"
	"masivo/interfaces"
	"masivo/model"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepo struct {
	interfaces.Repositorio
	db *sql.DB
}

func (repo *SqliteRepo) Nombre() string {
	return "SQLite"
}

func (repo *SqliteRepo) Inicializar() error {
	var err error
	repo.db, err = sql.Open("sqlite3", "./masivo.db")
	if err != nil {
		return err
	}
	return repo.db.Ping()
}

func (repo *SqliteRepo) InsertarLote(registros []model.Registro) error {
	values := []interface{}{}
	query := "INSERT INTO registros (nombre, domicilio, comentarios, puntaje, ingreso, fecha) VALUES "

	for _, registro := range registros {
		query += "(?, ?, ?, ?, ?, ?),"
		values = append(values, registro.Nombre, registro.Domicilio, registro.Comentarios, registro.Puntaje, registro.Ingreso, registro.Fecha)
	}

	query = query[:len(query)-1] // Remove the trailing comma

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	stmt, err = tx.Prepare("INSERT INTO registros (nombre, domicilio, comentarios, puntaje, ingreso, fecha) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, registro := range registros {
		_, err := stmt.Exec(registro.Nombre, registro.Domicilio, registro.Comentarios, registro.Puntaje, registro.Ingreso, registro.Fecha)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (repo *SqliteRepo) Limpiar() error {
	_, err := repo.db.Exec("DELETE FROM registros")
	return err
}

func (repo *SqliteRepo) Cerrar() error {
	if err := repo.db.Close(); err != nil {
		log.Println("Error closing the database connection:", err)
		return err
	}
	return nil
}
