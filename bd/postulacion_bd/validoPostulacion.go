package postulacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidoPostulacion(u postulacionmodels.Postulacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	registro := make(map[string]interface{})

	if u.Estado {
		registro["estado"] = u.Estado
		registro["mensaje"] = u.Mensaje
		registro["fechaPublicacion"] = u.FechaPublicacion
	}
	updtString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": u.ID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}
