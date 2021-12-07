package apimodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type TipoProyecto struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TipoProyecto   string         `bson:"tipoProyecto" json:"tipoProyecto,omitempty"`
	DecripcionTipo string         `bson:"descripcionTipo" json:"descripcionTipo,omitempty"`
	Presupuesto    float64        `bson:"presupuesto" json:"presupuesto,omitempty"`
}
