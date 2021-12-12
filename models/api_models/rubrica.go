package apimodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rubrica struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreRubrica      string             `bson:"nombreRubrica" json:"nombreRubrica,omitempty"`
	DescripcionRubrica string             `bson:"descripcionRubrica" json:"descripcionRubrica,omitempty"`
	PuntajeRubrica     float64            `bson:"puntajeRubrica" json:"puntajeRubrica,omitempty"`
}
