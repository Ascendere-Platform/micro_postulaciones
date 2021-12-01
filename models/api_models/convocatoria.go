package apimodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Convocatoria struct {
	ID                      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	NombreConvocatoria      string             `bson:"nombreConvocatoria" json:"nombreConvocatoria,omitempty"`
	PeriodoConvocatoria     string             `bson:"periodoConvocatoria" json:"periodoConvocatoria,omitempty"`
	CalificacionPostulacion float64            `bson:"calificacionPostulacion" json:"calificacionPostulacion,omitempty"`
	RubricasConvocatoria    []Rubrica          `bson:"rubricasConvocatoria" json:"rubricasConvocatoria,omitempty"`
}
