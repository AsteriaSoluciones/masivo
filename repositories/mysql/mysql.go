package mysql

import (
	"database/sql"
	"log"
	"masivo/model"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepo struct {
	db *sql.DB
}

func (repo *MysqlRepo) Nombre() string {
	return "MySQL"
}

func (repo *MysqlRepo) Inicializar() error {
	var err error
	repo.db, err = sql.Open("mysql", "masivo:Ap2485!aO65>AV_o8@tcp(127.0.0.1:3306)/masivo")
	if err != nil {
		return err
	}
	return repo.db.Ping()
}

func (repo *MysqlRepo) InsertarLote(registros []model.Registro) error {
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

func (repo *MysqlRepo) Limpiar() error {
	_, err := repo.db.Exec("TRUNCATE TABLE registros")
	return err
}

func (repo *MysqlRepo) Cerrar() error {
	if err := repo.db.Close(); err != nil {
		log.Println("Error closing the database connection:", err)
		return err
	}
	return nil
}
