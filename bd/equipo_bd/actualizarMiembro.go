package equipobd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ActualizoMiembro(id string, miembro apimodels.UsuarioEquipo) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	objID, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": objID}

	var resultado postulacionmodels.Postulacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return "No se encuentra la postulacion", err
	}

	for i := 0; i < len(resultado.Equipo); i++ {

		if resultado.Equipo[i].ID == miembro.ID {
			if resultado.Equipo[i].Cargo == "GESTOR" && miembro.Cargo != "GESTOR" {
				return "No puede modificar el rol del Gestor", err
			}
			if len(miembro.Asignatura) > 0 {
				resultado.Equipo[i].Asignatura = miembro.Asignatura
			}
			if len(miembro.Cargo) > 0 {
				resultado.Equipo[i].Cargo = miembro.Cargo
			}
		}
	}

	registro := make(map[string]interface{})

	registro["equipo"] = resultado.Equipo

	updtString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": resultado.ID}}

	_, errUpdt := col.UpdateOne(ctx, filtro, updtString)

	if errUpdt != nil {
		return "No se pude actualizar el miembro del equipo", err
	}

	return "", nil
}
