package postulacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoPostulaciones(tk string) ([]postulacionmodels.DevuelvoPostulacion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	var resultadoCompleto []postulacionmodels.DevuelvoPostulacion

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return resultadoCompleto, err
	}

	for cur.Next(ctx) {
		var postulacionSimple postulacionmodels.Postulacion
		err := cur.Decode(&postulacionSimple)
		if err != nil {
			return resultadoCompleto, err
		}

		postulacionCompleta, errorBusqueda := BuscoPostulacion(postulacionSimple.ID.Hex(), tk)
		if errorBusqueda != nil {
			return resultadoCompleto, errorBusqueda
		}

		resultadoCompleto = append(resultadoCompleto, postulacionCompleta)

	}

	err = cur.Err()
	if err != nil {
		return resultadoCompleto, err
	}
	cur.Close(ctx)
	return resultadoCompleto, nil
}
