package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Claim es la estructura usada para procesar JWT
type Claim struct {
	Email     string             `json:"email"`
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos string             `bson:"apellidos" json:"apellidos,omitempty"`
	Rol       string             `bson:"rol" json:"rol,omitempty"`
	jwt.StandardClaims
}
