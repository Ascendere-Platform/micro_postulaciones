package postulacionmodels

import (
	"time"

	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Postulacion struct {
	ID                 primitive.ObjectID        `bson:"_id,omitempty" json:"id"`
	NombreProyecto     string                    `bson:"nombreProyecto" json:"nombreProyecto,omitempty"`
	Justificacion      string                    `bson:"justificacion" json:"justificacion,omitempty"`
	Alcance            string                    `bson:"alcance" json:"alcance,omitempty"`
	Requerimientos     string                    `bson:"requerimientos" json:"requerimientos,omitempty"`
	Resultados         string                    `bson:"resultados" json:"resultados,omitempty"`
	CalificacionTotal  float64                   `bson:"calificacionTotal" json:"calificacionTotal,omitempty"`
	Restricciones      string                    `bson:"restricciones" json:"restricciones,omitempty"`
	Estado             bool                      `bson:"estado" json:"estado,omitempty"`
	FechaCreacion      time.Time                 `bson:"fechaInicio" json:"fechaInicio,omitempty"`
	FechaActualizacion time.Time                 `bson:"fechaFin" json:"fechaFin,omitempty"`
	Equipo             []apimodels.UsuarioEquipo `bson:"equipo" json:"equipo,omitempty"`
	ConvocatoriaID     primitive.ObjectID        `bson:"convocatoriaID" json:"convocatoriaID,omitempty"`
	Mensaje            string                    `bson:"mensaje" json:"mensaje,omitempty"`
	EvaluacionCompleta []primitive.ObjectID      `bson:"evaluacionCompleta" json:"evaluacionCompleta,omitempty"`
	TipoProyecto       primitive.ObjectID        `bson:"tipoProyectoId" json:"tipoProyectoId,omitempty"`
}
