package cronogramabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminoCronograma(tipoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("cronograma")

	objID, _ := primitive.ObjectIDFromHex(tipoID)

	condicion := bson.M{
		"_id": objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}