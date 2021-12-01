package cronogramabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	cronogramamodels "github.com/ascendere/micro-postulaciones/models/cronograma_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoCronograma(id string) (cronogramamodels.Cronograma, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("cronograma")

	objID, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": objID}

	var resultado cronogramamodels.Cronograma

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return resultado, err
	}
	return resultado, err
}