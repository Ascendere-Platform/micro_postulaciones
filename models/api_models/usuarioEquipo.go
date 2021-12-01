package apimodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsuarioEquipo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombres    string             `bson:"nombre" json:"nombre,omitempty"`
	Email      string             `bson:"email" json:"email"`
	Asignatura Asignatura         `bson:"asignatura" json:"asignatura"`
	Cargo      string             `bson:"cargo" json:"cargo"`
}
