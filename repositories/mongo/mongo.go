package mongo

import (
	"context"
	"masivo/interfaces"
	"masivo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	interfaces.Repositorio
	client *mongo.Client
}

func (repo *Mongo) Nombre() string {
	return "MongoDB"
}

func (repo *Mongo) Inicializar() error {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	newClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	repo.client = newClient

	// Check the connection
	err = repo.client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Mongo) InsertarLote(registros []model.Registro) error {
	collection := repo.client.Database("masiva").Collection("registros")

	_, err := collection.InsertMany(context.TODO(), convertRegistrosToRows(registros))
	if err != nil {
		return err
	}
	return nil
}

func (repo *Mongo) Limpiar() error {
	collection := repo.client.Database("masiva").Collection("registros")

	err := collection.Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (repo *Mongo) Cerrar() error {
	err := repo.client.Disconnect(context.Background())
	return err
}

// convertRegistrosToRows convierte un slice de registros a un slice de filas
// para ser insertadas en la base de datos
func convertRegistrosToRows(registros []model.Registro) []interface{} {
	var rows []interface{}
	for _, registro := range registros {
		row := bson.D{
			{Key: "nombre", Value: registro.Nombre},
			{Key: "domicilio", Value: registro.Domicilio},
			{Key: "comentarios", Value: registro.Comentarios},
			{Key: "puntaje", Value: registro.Puntaje},
			{Key: "ingreso", Value: registro.Ingreso},
			{Key: "fecha", Value: registro.Fecha},
		}
		rows = append(rows, row)
	}

	return rows
}
