package evaluacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	evaluacionmodels "github.com/ascendere/micro-postulaciones/models/evaluacion_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoEvaluacion(id string) (evaluacionmodels.EvaluacionPostulacion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("evaluaciones")

	objID, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": objID}

	var resultado evaluacionmodels.EvaluacionPostulacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return resultado, err
	}
	return resultado, err
}