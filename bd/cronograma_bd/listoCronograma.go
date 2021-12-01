package cronogramabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	cronogramamodels "github.com/ascendere/micro-postulaciones/models/cronograma_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoCronograma() ([]*cronogramamodels.Cronograma, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("cronograma")

	var results []*cronogramamodels.Cronograma

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, false
	}

	for cur.Next(ctx) {
		var s cronogramamodels.Cronograma
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}
		results = append(results, &s)

	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true

}