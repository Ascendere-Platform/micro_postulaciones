package postulacionbd

import (
	"context"
	"time"

	"github.com/ascendere/micro-postulaciones/bd"
	apibd "github.com/ascendere/micro-postulaciones/bd/api_bd"
	evaluacionbd "github.com/ascendere/micro-postulaciones/bd/evaluacion_bd"
	postulacionmodels "github.com/ascendere/micro-postulaciones/models/postulacion_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPostulacion(id string, tk string) (postulacionmodels.DevuelvoPostulacion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Postualciones")
	col := db.Collection("propuestas")

	objID, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": objID}

	var resultado postulacionmodels.Postulacion
	var propuestaEncontrada postulacionmodels.DevuelvoPostulacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return propuestaEncontrada, err
	}

	propuestaEncontrada.ID = resultado.ID
	propuestaEncontrada.NombreProyecto = resultado.NombreProyecto
	propuestaEncontrada.Alcance = resultado.Alcance
	propuestaEncontrada.Justificacion = resultado.Justificacion
	propuestaEncontrada.Requerimientos = resultado.Requerimientos
	propuestaEncontrada.Resultados = resultado.Resultados
	propuestaEncontrada.CalificacionTotal = resultado.CalificacionTotal
	propuestaEncontrada.Restricciones = resultado.Restricciones
	propuestaEncontrada.Estado = false
	propuestaEncontrada.FechaCreacion = resultado.FechaCreacion
	propuestaEncontrada.FechaActualizacion = resultado.FechaActualizacion
	idConvocatoria := resultado.ConvocatoriaID.Hex()
	propuestaEncontrada.Convocatoria,_ = apibd.ValidoConvocatoria(idConvocatoria, tk)
	propuestaEncontrada.Mensaje = resultado.Mensaje
	propuestaEncontrada.TipoProyecto,_ = apibd.ValidoTipo(resultado.TipoProyecto.Hex(),tk)

	for _, miembro := range resultado.Equipo {
		miembroEncontrado, _:= apibd.DevuelvoUsuarioEquipo(miembro, tk)
		propuestaEncontrada.Equipo = append(propuestaEncontrada.Equipo, miembroEncontrado)
	}
	
	for _, evaluacionID := range resultado.EvaluacionCompleta {
		evaluacionCompleta,errBusqueda := evaluacionbd.BuscoEvaluacion(evaluacionID.Hex())

		if errBusqueda != nil {
			propuestaEncontrada.EvaluacionCompleta = append(propuestaEncontrada.EvaluacionCompleta, evaluacionCompleta)
		}
	}

	return propuestaEncontrada, err

}
