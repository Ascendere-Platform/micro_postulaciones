package apimodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Asignatura struct {
		ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		NombreAsignatura string             `bson:"nombreAsignatura" json:"nombreAsignatura,omitempty"`
		Modalidad        string             `bson:"modalidad" json:"modalidad,omitempty"`
		FacultadID       string             `bson:"facultadid" json:"facultadid,omitempty"`
		Periodo          string             `bson:"periodo" json:"periodo,omitempty"`
}