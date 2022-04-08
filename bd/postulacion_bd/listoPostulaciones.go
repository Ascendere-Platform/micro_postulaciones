package postulacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoPostulaciones(tk string) ([]postulacionmodels.DevuelvoListado, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	var results []postulacionmodels.DevuelvoListado

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, err
	}

	for cur.Next(ctx) {
		var s postulacionmodels.DevuelvoListado
		err := cur.Decode(&s)
		if err != nil {
			return results, err
		}
			results = append(results, s)

	}

	err = cur.Err()
	if err != nil {
		return results, nil
	}
	cur.Close(ctx)
	return results, err

}