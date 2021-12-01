package cronogramamodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cronograma struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreHito    string             `bson:"nombreHito" json:"nombreHito,omitempty"`
	Entregables          []string             `bson:"entregables" json:"entregables,omitempty"`
	FechaInicio time.Time          `bson:"fechaInicio" json:"fechaInicio,omitempty"`
	FechaFin time.Time          `bson:"fechaFin" json:"fechaFin,omitempty"`
}
