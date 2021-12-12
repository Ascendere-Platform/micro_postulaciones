package postulacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ActualizoPostulacion(u postulacionmodels.Postulacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	registro := make(map[string]interface{})

	if len(u.NombreProyecto) > 0 {
		registro["nombreProyecto"] = u.NombreProyecto
	}

	if len(u.Alcance) > 0 {
		registro["periodoConvocatoria"] = u.Alcance
	}

	if len(u.ConvocatoriaID) > 0 {
		registro["antecedentes"] = u.ConvocatoriaID
	}

	if len(u.Restricciones) > 0 {
		registro["objetivos"] = u.Restricciones
	}

	if len(u.TipoProyecto) > 0 {
		registro["banner"] = u.TipoProyecto
	}

	if len(u.Justificacion) > 0 {
		registro["destinatario"] = u.Justificacion
	}

	if len(u.Resultados) > 0 {
		registro["reconocimiento"] = u.Resultados
	}

	if len(u.Restricciones) > 0 {
		registro["compromisos"] = u.Restricciones
	}

	registro["fechaActualizacion"] = u.FechaActualizacion
	registro["mensaje"]=u.Mensaje

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
