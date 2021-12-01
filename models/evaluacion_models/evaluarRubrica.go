package convocatoriamodels

import (
	apimodels "github.com/ascendere/micro-postulaciones/models/api_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EvaluacionRubrica struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Rubrica      apimodels.Rubrica  `bson:"rubrica" json:"rubrica,omitempty"`
	Calificacion float64            `bson:"calificacion" json:"calificacion,omitempty"`
	Mensaje      string             `bson:"mensaje" json:"mensaje,omitempty"`
	Estado       bool               `bson:"estado" json:"estado,omitempty"`
}
