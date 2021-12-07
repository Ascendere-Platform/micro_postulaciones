package postulacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	apibd "github.com/ascendere/micro-postulaciones/bd/api_bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistrarPostulacion(r postulacionmodels.Postulacion, tk string) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	registro := postulacionmodels.Postulacion{
		ID:             primitive.NewObjectID(),
		NombreProyecto: r.NombreProyecto,
		Justificacion:  r.Justificacion,
		Alcance:        r.Alcance,
		Requerimientos: r.Requerimientos,
		Resultados:     r.Resultados,
		Restricciones:  r.Restricciones,
		Estado:         r.Estado,
		FechaCreacion:  r.FechaCreacion,
		ConvocatoriaID: r.ConvocatoriaID,
		TipoProyecto:   r.TipoProyecto,
		Mensaje: r.Mensaje,
	}

	for _, miembro := range r.Equipo {
		personal, _ := apibd.ValidoUsuario(miembro.ID.Hex(), tk)
		miembro.Email = personal.Email
		miembro.Nombres = personal.Nombres
		registro.Equipo = append(registro.Equipo, miembro)
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
