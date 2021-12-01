package cronogramabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	cronogramamodels "github.com/ascendere/micro-postulaciones/models/cronograma_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistrarCronograma(r cronogramamodels.Cronograma) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("cronograma")

	registro := cronogramamodels.Cronograma{
		ID:                 primitive.NewObjectID(),
		NombreHito: r.NombreHito,
		FechaInicio: r.FechaInicio,
		FechaFin: r.FechaFin,
		Entregables: r.Entregables,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}