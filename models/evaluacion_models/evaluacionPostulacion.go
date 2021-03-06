package evaluacionmodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type EvaluacionPostulacion struct {
	ID                 primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	RubricasEvaluadas []EvaluacionRubrica `bson:"rubricaEvaluada" json:"rubricaEvaluada,omitempty"`
	Calificacion float64 `bson:"calificacion" json:"calificacion,omitempty"`
	Evaluador struct {
		Email  string             `json:"email"`
		ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		Nombre string             `bson:"nombre" json:"nombre,omitempty"`
	}
}